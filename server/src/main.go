package main

import (
	"database/sql"

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
	server.Run()
}
