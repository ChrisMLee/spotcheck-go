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

// graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *sql.DB) *Root {
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
						Type: User,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							// Strip the name from arguments and assert that it's a string
							userId, ok := p.Args["id"].(int)
							fmt.Println("trying")
							if ok {
								var uid int
								var un string
								var ue string
								sqlStatement := `SELECT id, username, email FROM users WHERE id=$1`
								row := db.QueryRow(sqlStatement, userId)
								err := row.Scan(&uid, &un, &ue)
								if err != nil {
									if err == sql.ErrNoRows {
										fmt.Println("Zero rows found")
										return nil, err
									} else {
										panic(err)
									}
								}

								user := userData{Id: uid, Username: un, Email: ue}
								return user, nil
							}
							return nil, nil
						},
					},
				},
			},
		),
	}
	return &root
}
