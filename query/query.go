package query

import (
	"errors"
	"graphql-example/object"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getPost": &graphql.Field{
				Type: object.PostType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					for _, post := range object.Posts {
						if post.ID == id {
							return post, nil
						}
					}
					return nil, errors.New("post not found")
				},
			},
			"getPosts": &graphql.Field{
				Type: graphql.NewList(object.PostType),
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					return object.Posts, nil
				},
			},
		},
	},
)
