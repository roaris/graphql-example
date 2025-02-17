package mutation

import (
	"errors"
	"graphql-example/object"
	"graphql-example/util"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type: object.PostType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userID := util.GetUserIDFromContext(p.Context)
					post := object.Post{
						ID:      uuid.New().String(),
						Title:   p.Args["title"].(string),
						Content: p.Args["content"].(string),
						UserID:  userID,
					}
					object.Posts = append(object.Posts, post)
					return post, nil
				},
			},
			"updatePost": &graphql.Field{
				Type: object.PostType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userID := util.GetUserIDFromContext(p.Context)
					postID := p.Args["id"].(string)
					title := p.Args["title"].(string)
					content := p.Args["content"].(string)
					for i, post := range object.Posts {
						if post.ID == postID {
							if post.UserID != userID {
								return nil, errors.New("forbidden")
							}
							object.Posts[i].Title = title
							object.Posts[i].Content = content
							return object.Posts[i], nil
						}
					}
					return nil, errors.New("post not found")
				},
			},
			"deletePost": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userID := util.GetUserIDFromContext(p.Context)
					postID := p.Args["id"].(string)
					for i, post := range object.Posts {
						if post.ID == postID {
							if post.UserID != userID {
								return nil, errors.New("forbidden")
							}
							object.Posts = append(object.Posts[:i], object.Posts[i+1:]...)
							return postID, nil
						}
					}
					return nil, errors.New("post not found")
				},
			},
		},
	},
)
