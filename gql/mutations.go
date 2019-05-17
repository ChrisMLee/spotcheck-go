package gql

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

var CreateSpotType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateSpot",
	Fields: graphql.InputObjectConfigFieldMap{
		"lat": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The latitude coordinate of the spot",
		},
		"lng": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The longitude coordinate of the spot",
		},
		"image_url": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "A link to an image of the spot",
		},
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the spot",
		},
		"description": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "A description of the spot",
		},
		"address": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "The address of the spot",
		},
		"user_id": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The id of the creating user",
		},
	},
})

var Mutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "MutationType",
	Fields: graphql.Fields{
		"createSpot": &graphql.Field{
			Type: SpotType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{
					Description: "An input with the spot details",
					Type:        graphql.NewNonNull(CreateSpotType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				inp, ok := p.Args["input"].(map[string]interface{})
				if ok {
					sqlStatement := `INSERT INTO spots (name, image_url, description, address, lat, lng, user_id) VALUES (
						$1,
						$2,
						$3,
						$4,
						$5,
						$6,
						$7
					) RETURNING id;`

					db, _ := p.Context.Value("db").(*sql.DB)
					name := NewNullString(inp["name"].(string))
					imageUrl := NewNullString(inp["image_url"].(string))
					description := NewNullString(inp["description"].(string))
					address := NewNullString(inp["address"].(string))
					lat := inp["lat"].(string)
					lng := inp["lng"].(string)
					userId := inp["user_id"].(int)

					newSpotId := 0
					err := db.QueryRow(sqlStatement, name, imageUrl, description, address, lat, lng, userId).Scan(&newSpotId)

					if err != nil {
						fmt.Println("GOT AN ERROR")
						fmt.Println(err)
						panic(err)
					}
					fmt.Println("newSpotId")
					fmt.Println(newSpotId)
					newSpotSqlStatement := `SELECT id, name, image_url, description, address, lat, lng, user_id FROM spots WHERE id=$1`
					row := db.QueryRow(newSpotSqlStatement, newSpotId)

					var dbSpotId int
					var dbSpotName NullString
					var dbSpotImageUrl NullString
					var dbSpotDescription NullString
					var dbSpotAddress NullString
					var dbSpotLat string
					var dbSpotLng string
					var dbSpotUserId int

					err = row.Scan(
						&dbSpotId,
						&dbSpotName,
						&dbSpotImageUrl,
						&dbSpotDescription,
						&dbSpotAddress,
						&dbSpotLat,
						&dbSpotLng,
						&dbSpotUserId,
					)
					if err != nil {
						fmt.Println(err)
						if err == sql.ErrNoRows {
							fmt.Println("Zero rows found")
							return nil, err
						} else {
							panic(err)
						}
					}

					newSpot := spot{Id: dbSpotId, Name: dbSpotName, ImageUrl: dbSpotImageUrl, Address: dbSpotAddress, Lat: dbSpotLat, Lng: dbSpotLng, UserId: dbSpotUserId}
					fmt.Println("newSpot")
					fmt.Println(newSpot)
					return newSpot, nil

				}
				return nil, nil

			},
		},
	},
})
