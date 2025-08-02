package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"ToDoServer/tododatabase"
	. "ToDoServer/datatypes"

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

	// Run web server
	fs := http.FileServer(http.Dir("../../client/dist"))
	http.Handle("/", fs)
	http.HandleFunc("GET /api/todos", todosHandler)
	http.HandleFunc("POST /api/newtodo", newTodoHandler)

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

func todosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := tododatabase.GetTodos(db)
	if err != nil {
		log.Fatal("Unable to get todos from database")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Unable to marshal todos into json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	w.Write(result)
}

func newTodoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo Todo
	err := decoder.Decode(&todo)
	if err != nil {
		log.Fatal("Unable to decode request into struct")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tododatabase.AddTodo(db, &todo)
	if err != nil {
		log.Fatal("Unable to add new todo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	result, err := json.Marshal(todo)
	if err != nil {
		log.Fatal("Unable to marshal newly created todo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Controll-Allow-Origin", "http://localhost:5173")

	w.Write(result)
}