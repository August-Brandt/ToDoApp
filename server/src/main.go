package main

import (
	"database/sql"
	"log"

	"ToDoServer/todoServer"
	"ToDoServer/tododatabase"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// Setup sqlite database
	var err error
	db, err = tododatabase.Setup("sqlite3", "../database/ToDoDatabase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := todoServer.NewServer(":8080", db, "../../client/dist")
	if err := server.Run(); err != nil {
		log.Printf("Server stopped: %s\n", err.Error())
	}
}
