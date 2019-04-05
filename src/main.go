package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

var db *sql.DB

// https://github.com/gin-gonic/examples/blob/master/basic/main.go

// https://github.com/gin-gonic/gin/blob/master/context.go

// https://godoc.org/github.com/gin-gonic/gin

// func make(t Type, size ...IntegerType) Type
var db = make(map[string]string)

func main() {
	initDb()
	r := gin.Default()
	db, err := sql.Open("postgres", "user=spotcheck-db-dev dbname=spotcheck_dev sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:id", func(c *gin.Context) {
		fmt.Println("tried to print")
		id := c.Param("id")
		value, ok := db[id]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": id, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": id, "status": "no value"})
		}
	})

	r.POST("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		var json struct {
			Value string `json: "value" binding:"required"`
		}
		if c.Bind(&json) == nil {
			db[id] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func initDb() {
	fmt.Println("Trying to init the db")
	config := dbConfig()
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
