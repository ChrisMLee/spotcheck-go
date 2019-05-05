package gql

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

type userData struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// type Hero struct {
// 	Id      string `graphql:"id"`
// 	Name    string
// 	Friends []Hero `graphql:"friends"`
// }

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *sql.DB) *Root {
	resolver := Resolver{db: db}
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQuery",
				Fields: graphql.Fields{
					"hello": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "world", nil
						},
					},
					"ron": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "artest", nil
						},
					},
					"user": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: UserType,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.UserResolver},
				},
			},
		),
	}
	return &root
}
