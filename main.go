package main

import (
	"database/sql"
	"log"
	"net/http"
	//_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	//initialize databse
	initDB()
	//start server
	log.Fatal(http.ListenAndServe(":4200", nil))
}

func initDB() {
	var err error
	//connecting to postgresql
	db, err = sql.Open("postgres", "dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
}
