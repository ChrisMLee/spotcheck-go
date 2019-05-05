package gql

import "github.com/graphql-go/graphql"

type spot struct {
	Name        NullString `json:"name"`
	ImageUrl    NullString `json:"image_url"`
	Description NullString `json:"description"`
	Address     NullString `json:"address"`
	Lat         string     `json:"lat"`
	Lng         string     `json:"lng"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			// "spots": &graphql.Field{
			// 	Type:        graphql.NewList(SpotType),
			// 	Description: "A list of spots created by a given user",
			// 	Resolver: func(p graphql.ResolveParams) (interface{}, error) {
			// 		userId, ok := p.Args["id"].(int)
			// 		spots := make([]spot, 0)
			// 		sqlStatement := `SELECT name, image_url, description, address, lat, lng FROM spots WHERE user_id=$1`
			// 		rows, err := db.Query(sqlStatement, userId)
			// 		defer rows.Close()
			// 		for rows.Next() {
			// 			spot := spot{}
			// 			if err := rows.Scan(
			// 				&spot.Name,
			// 				&spot.ImageUrl,
			// 				&spot.Description,
			// 				&spot.Address,
			// 				&spot.Lat,
			// 				&spot.Lng); err != nil {
			// 					log.Println(err)
			// 				}
			// 			log.Println(spot)
			// 			spots = append(spots, spot)
			// 		}
			// 		err = rows.Err()
			// 		if err != nil {
			// 			c.JSON(500, err)
			// 			return
			// 		}

			// 		c.JSON(200, spotResponse{Spots: spots})

			//         },
			// },
		},
	},
)

var SpotType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Spot",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"lat": &graphql.Field{
				Type: graphql.String,
			},
			"lng": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
