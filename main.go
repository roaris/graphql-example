package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"graphql-example/mutation"
	"graphql-example/object"
	"graphql-example/query"
	"graphql-example/util"
	"log"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

type loginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type graphqlReq struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

var schema graphql.Schema

func init() {
	var err error
	if schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.QueryType,
			Mutation: mutation.MutationType,
		},
	); err != nil {
		log.Fatalln(err)
	}

	// disable introspection (ref: https://github.com/graphql-go/graphql/issues/649)
	graphql.SchemaMetaFieldDef.Resolve = func(p graphql.ResolveParams) (interface{}, error) {
		return nil, errors.New("introspection is disabled")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" { // this validation prevents csrf by text/plain
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, user := range object.Users {
		if user.Name == req.Name && user.Password == req.Password {
			token := util.GenerateJWT(user.ID)
			w.Header().Set("Set-Cookie", fmt.Sprintf("session=%s", token))
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "invalid credentials")
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" { // this validation prevents csrf by text/plain
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req graphqlReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := r.Cookie("session")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	userID, err := util.VerifyJWT(token.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx := util.SetUserIDToContext(r.Context(), userID)

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  req.Query,
		OperationName:  req.OperationName,
		VariableValues: req.Variables,
		Context:        ctx,
	})
	// disable suggestion
	for i := range result.Errors {
		p := strings.Index(result.Errors[i].Message, " Did you mean")
		if p != -1 {
			result.Errors[i].Message = result.Errors[i].Message[:p]
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/graphql", graphqlHandler)

	log.Println("Server is running on port 8000")
	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		log.Fatalln(err)
	}
}
