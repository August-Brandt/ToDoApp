package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	dt "ToDoServer/datatypes"
	"ToDoServer/tododatabase"

	"github.com/google/uuid"
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

	// Setup endpoints server
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../../client/dist"))
	mux.Handle("/", fs)
	mux.HandleFunc("/api/newtodo", newTodoHandler)
	mux.HandleFunc("/api/removetodo", removeTodoHandler)
	mux.HandleFunc("/api/finishtodo", finishTodoHandler)
	mux.HandleFunc("/api/unfinishtodo", UnfinishTodoHandle)
	mux.HandleFunc("GET /api/todos", todosHandler)

	// Start webserver
	port, exists := os.LookupEnv("PORT")
	if exists {
		port = ":" + port
	} else {
		port = ":8080"
	}
	fmt.Printf("Website can be found at http://localhost%s\n", port)

	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		log.Fatal("Invalid method used for /api/newtodo")
		http.Error(w, "Invalid method used", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	todo := &dt.Todo{}

	err := decoder.Decode(todo)
	if err != nil {
		log.Fatal("Unable to decode request into struct")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Id = uuid.NewString()
	id, err := tododatabase.AddTodo(db, todo)
	if err != nil {
		log.Fatalf("Unable to add new todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTodo, err := tododatabase.GetTodoById(db, id)
	if err != nil {
		log.Fatalf("Unable to retreve newly added todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(newTodo)
	if err != nil {
		log.Fatal("Unable to marshal newly created todo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Added new todo to database")
	w.Write(result)
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		log.Fatal("Invalid method used for /api/removetodo")
		http.Error(w, "Invalid method used", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	todo := &dt.Todo{}

	err := decoder.Decode(todo)
	if err != nil {
		log.Fatalf("Unable to decode request body\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = tododatabase.RemoveTodoById(db, todo.Id)
	if err != nil {
		log.Fatalf("Unable to remove todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Removed todo from database")
	w.Write([]byte{})
}

func finishTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		log.Fatal("Invalid method used for /api/finishtodo")
		http.Error(w, "Invalid method used", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	todo := &dt.Todo{}

	err := decoder.Decode(todo)
	if err != nil {
		log.Fatalf("Unable to decode request body\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = tododatabase.FinishTodo(db, todo.Id)
	if err != nil {
		log.Fatalf("Unable to update todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Updated a todo to be finished")
	w.Write([]byte{})
}

func UnfinishTodoHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		log.Fatal("Invalid method used for /api/unfinishtodo")
		http.Error(w, "Invalid method used", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	todo := &dt.Todo{}

	err := decoder.Decode(todo)
	if err != nil {
		log.Fatalf("Unable to decode request body\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = tododatabase.UnfinishTodo(db, todo.Id)
	if err != nil {
		log.Fatalf("Unable to unfinish todo in database\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Updated a todo to be unfinished")
	w.Write([]byte{})
}