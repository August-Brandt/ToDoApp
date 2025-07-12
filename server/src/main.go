package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ToDoServer/tododatabase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup sqlite database
	db, err := tododatabase.Setup("sqlite3", "../database/ToDoDatabase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Run web server
	fs := http.FileServer(http.Dir("../../client/dist"))
	http.Handle("/", fs)
	port, exists := os.LookupEnv("PORT")
	if exists {
		port = ":" + port
	} else {
		port = ":8080"
	}
	fmt.Printf("Website can be found at http://localhost%s\n", port)

	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
