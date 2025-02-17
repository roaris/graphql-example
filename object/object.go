package object

import (
	"errors"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string
}

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string
}

var Users = []User{
	{ID: "04b4d343-d974-451e-bb14-cc5755204c7a", Name: "user1", Password: "password1"},
	{ID: "67ed72e2-7ffb-4427-9450-b87aab994703", Name: "user2", Password: "password2"},
}

var Posts = []Post{
	{ID: "b711966f-1a9d-46d5-8f5e-c2bfe8f94229", Title: "title1", Content: "content1", UserID: Users[0].ID},
	{ID: "0ae944de-40a9-447d-9066-89d3a190e32f", Title: "title2", Content: "content2", UserID: Users[0].ID},
	{ID: "ac95a094-82ba-47ed-ae84-f46053816544", Title: "title3", Content: "content3", UserID: Users[1].ID},
	{ID: "526875df-18f3-44ae-b60e-908a935e605f", Title: "title4", Content: "content4", UserID: Users[1].ID},
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: UserType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post := p.Source.(Post)
					for _, user := range Users {
						if user.ID == post.UserID {
							return user, nil
						}
					}
					return nil, errors.New("user not found")
				},
			},
		},
	},
)

// fix initialization cycle (ref: https://github.com/graphql-go/graphql/issues/164)
func init() {
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewList(PostType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := p.Source.(User)
			var userPosts []Post
			for _, post := range Posts {
				if user.ID == post.UserID {
					userPosts = append(userPosts, post)
				}
			}
			return userPosts, nil
		},
	})
}
