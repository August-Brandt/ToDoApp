package todoServer

import (
	dt "ToDoServer/datatypes"
	"ToDoServer/tododatabase"
	"fmt"

	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *ToDoServer) SetRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.FileServer(http.Dir(s.frontEnd)))

	middlewareChain := []Middleware{
		loggingMiddleware,
	}

	if s.Mode == "DEV" {
		middlewareChain = append(middlewareChain, corsMiddleware)
	}

	fmt.Println(middlewareChain)

	// API handlers
	mux.HandleFunc("GET /api/todos", ApplyMiddlewareChain(s.todosHandler, middlewareChain...))
	mux.HandleFunc("POST /api/addtodo", ApplyMiddlewareChain(s.newTodoHandler, middlewareChain...))
	mux.HandleFunc("POST /api/removetodo/{todoID}", ApplyMiddlewareChain(s.removeTodoHandler, middlewareChain...))
	mux.HandleFunc("POST /api/finishtodo/{todoID}", ApplyMiddlewareChain(s.finishTodoHandler, middlewareChain...))
	mux.HandleFunc("POST /api/unfinishtodo/{todoID}", ApplyMiddlewareChain(s.unfinishTodoHandle, middlewareChain...))
}

func (s *ToDoServer) todosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := tododatabase.GetTodos(s.DB)
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

	w.Write(result)
}

func (s *ToDoServer) newTodoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	todo := &dt.Todo{}

	err := decoder.Decode(todo)
	if err != nil {
		log.Fatal("Unable to decode request into struct")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.Id = uuid.NewString()
	id, err := tododatabase.AddTodo(s.DB, todo)
	if err != nil {
		log.Fatalf("Unable to add new todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTodo, err := tododatabase.GetTodoById(s.DB, id)
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

func (s *ToDoServer) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("todoID")

	err := tododatabase.RemoveTodoById(s.DB, id)
	if err != nil {
		log.Fatalf("Unable to remove todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Removed todo from database")
	w.WriteHeader(http.StatusOK)
}

func (s *ToDoServer) finishTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("todoID")
	err := tododatabase.FinishTodo(s.DB, id)
	if err != nil {
		log.Fatalf("Unable to update todo\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Updated a todo to be finished")
	w.WriteHeader(http.StatusOK)
}

func (s *ToDoServer) unfinishTodoHandle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("todoID")
	err := tododatabase.UnfinishTodo(s.DB, id)
	if err != nil {
		log.Fatalf("Unable to unfinish todo in database\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Updated a todo to be unfinished")
	w.WriteHeader(http.StatusOK)
}
