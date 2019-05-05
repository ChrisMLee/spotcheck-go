package gql

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *sql.DB
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	userId, ok := p.Args["id"].(int)
	if ok {
		var uid int
		var un string
		var ue string
		sqlStatement := `SELECT id, username, email FROM users WHERE id=$1`
		row := r.db.QueryRow(sqlStatement, userId)
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
}
