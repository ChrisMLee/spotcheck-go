package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

var db *sql.DB

// https://github.com/gin-gonic/examples/blob/master/basic/main.go

// https://github.com/gin-gonic/gin/blob/master/context.go

// https://godoc.org/github.com/gin-gonic/gin

// func make(t Type, size ...IntegerType) Type
var fakeDb = make(map[string]string)

func main() {
	initDb()
	r := gin.Default()
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

		out, err := json.Marshal(userData{Id: uid, Username: un, Email: ue})
		if err != nil {
			c.JSON(500, err)
			return
		}

		c.JSON(200, gin.H{
			"message": string(out),
		})
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
