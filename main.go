package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/chrismlee/spotcheck-go/gql"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	dbhost = "DB_HOST"
	dbport = "DB_PORT"
	dbuser = "DB_USER"
	dbpass = "DB_PASSWORD"
	dbname = "DB_NAME"
)

type userData struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type spot struct {
	Name        NullString `json:"name"`
	ImageUrl    NullString `json:"image_url"`
	Description NullString `json:"description"`
	Address     NullString `json:"address"`
	Lat         string     `json:"lat"`
	Lng         string     `json:"lng"`
}

type spotResponse struct {
	Spots []spot `json:"spots"`
}

type reqBody struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// NullString is an alias for sql.NullString data type
// via https://gist.github.com/rsudip90/022c4ef5d98130a224c9239e0a1ab397
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

var db *sql.DB

// https://github.com/gin-gonic/examples/blob/master/basic/main.go

// https://github.com/gin-gonic/gin/blob/master/context.go

// https://godoc.org/github.com/gin-gonic/gin

// func make(t Type, size ...IntegerType) Type
var fakeDb = make(map[string]string)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	initDb()
	r := gin.Default()
	r.Use(CORSMiddleware())

	db, err := sql.Open("postgres", "user=spotcheck_db_dev dbname=spotcheck_dev sslmode=disable")
	if err != nil {
		panic(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("da ping")
		// data := userData{}

		userId := 1
		var uid int
		var un string
		var ue string
		sqlStatement := `SELECT id, username, email FROM users WHERE id=$1`
		row := db.QueryRow(sqlStatement, userId)
		err := row.Scan(&uid, &un, &ue)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err)
			}
		}

		out, err := json.Marshal(userData{Id: uid, Username: un, Email: ue})
		fmt.Println(un)
		if err != nil {
			c.JSON(500, err)
			return
		}
		c.JSON(200, gin.H{
			"message": string(out),
		})
	})
	r.GET("/user/:id", func(c *gin.Context) {
		// value, ok := fakeDb[id]
		// if ok {
		// 	c.JSON(http.StatusOK, gin.H{"user": id, "value": value})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{"user": id, "status": "no value"})
		// }

		userId := c.Param("id")
		var uid int
		var un string
		var ue string
		sqlStatement := `SELECT id, username, email FROM users WHERE id=$1`
		row := db.QueryRow(sqlStatement, userId)
		err := row.Scan(&uid, &un, &ue)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
				c.JSON(http.StatusOK, gin.H{"user": userId, "status": "no value"})
				return
			} else {
				panic(err)
			}
		}

		c.JSON(200, userData{Id: uid, Username: un, Email: ue})
	})

	r.GET("/user/:id/spots", func(c *gin.Context) {
		userId := c.Param("id")
		spots := make([]spot, 0)
		sqlStatement := `SELECT name, image_url, description, address, lat, lng FROM spots WHERE user_id=$1`
		rows, err := db.Query(sqlStatement, userId)
		defer rows.Close()
		for rows.Next() {
			spot := spot{}
			if err := rows.Scan(
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
			c.JSON(500, err)
			return
		}

		c.JSON(200, spotResponse{Spots: spots})
	})

	r.POST("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		var json struct {
			Value string `json: "value" binding:"required"`
		}
		if c.Bind(&json) == nil {
			fakeDb[id] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	rootQuery := gql.NewRoot(db)
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query, Mutation: gql.Mutations},
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	r.POST("/graphql", func(c *gin.Context) {

		var rBody reqBody

		c.BindJSON(&rBody)
		fmt.Println("the body")
		fmt.Println(rBody.Query)
		fmt.Println("the variables")
		fmt.Println(rBody.Variables)

		ctx := context.WithValue(context.Background(), "db", db)
		params := graphql.Params{Schema: schema, RequestString: rBody.Query, VariableValues: rBody.Variables, Context: ctx}
		r := graphql.Do(params)

		if len(r.Errors) > 0 {
			x, _ := ioutil.ReadAll(c.Request.Body)
			fmt.Printf("%s", string(x))

		}

		c.JSON(200, r)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func initDb() {
	fmt.Println("Trying to init the db")
	// config := dbConfig()
}

func dbConfig() map[string]string {
	fmt.Println("running the dbConfig")
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	// port, ok := os.LookupEnv(dbport)
	// if !ok {
	// 	panic("DBPORT environment variable required but not set")
	// }
	// user, ok := os.LookupEnv(dbuser)
	// if !ok {
	// 	panic("DBUSER environment variable required but not set")
	// }
	// password, ok := os.LookupEnv(dbpass)
	// if !ok {
	// 	panic("DBPASS environment variable required but not set")
	// }
	// name, ok := os.LookupEnv(dbname)
	// if !ok {
	// 	panic("DBNAME environment variable required but not set")
	// }
	// conf[dbhost] = host
	// conf[dbport] = port
	// conf[dbuser] = user
	// conf[dbpass] = password
	// conf[dbname] = name
	conf[dbhost] = host
	return conf
}
