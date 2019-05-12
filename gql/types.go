package gql

import (
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
)

// type spot struct {
// 	Name        NullString `json:"name"`
// 	ImageUrl    NullString `json:"image_url"`
// 	Description NullString `json:"description"`
// 	Address     NullString `json:"address"`
// 	Lat         string     `json:"lat"`
// 	Lng         string     `json:"lng"`
// }

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(userData)
				return user.Id, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(userData)
				fmt.Println("damn son where'd you find this?")
				fmt.Println("email")
				fmt.Println(user)
				return user.Email, nil
			},
		},
		"username": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(userData)
				fmt.Println("username")
				fmt.Println(user)
				return user.Username, nil
			},
		},
		"spots": &graphql.Field{
			Type: graphql.NewList(SpotType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := p.Source.(userData)

				spots := make([]spot, 0)
				sqlStatement := `SELECT id, name, image_url, description, address, lat, lng FROM spots WHERE user_id=$1`
				db, _ := p.Context.Value("db").(*sql.DB)
				rows, err := db.Query(sqlStatement, user.Id)
				defer rows.Close()
				for rows.Next() {
					spot := spot{}
					if err := rows.Scan(
						&spot.Id,
						&spot.Name,
						&spot.ImageUrl,
						&spot.Description,
						&spot.Address,
						&spot.Lat,
						&spot.Lng); err != nil {
						log.Println(err)
					}
					log.Println(spot)
					spots = append(spots, spot)
				}
				err = rows.Err()
				if err != nil {
					if err == sql.ErrNoRows {
						fmt.Println("Zero rows found")
						return nil, err
					} else {
						panic(err)
					}
				}
				return spots, nil
			},
		},
	},
})

var SpotType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Spot",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					spot := p.Source.(spot)
					return spot.Id, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					spot := p.Source.(spot)
					if spot.Name.Valid == true {
						return spot.Name.String, nil
					} else {
						return nil, nil
					}
				},
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					spot := p.Source.(spot)
					if spot.ImageUrl.Valid == true {
						return spot.ImageUrl.String, nil
					} else {
						return nil, nil
					}
				},
			},
			"description": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					spot := p.Source.(spot)
					if spot.Description.Valid == true {
						return spot.Description.String, nil
					} else {
						return nil, nil
					}
				},
			},
			"address": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					spot := p.Source.(spot)
					if spot.Address.Valid == true {
						return spot.Address.String, nil
					} else {
						return nil, nil
					}
				},
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
