package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Dodate      string
	Finished    bool
}

func main() {
	// Setup sqlite database
	db, err := sql.Open("sqlite3", "../database/ToDoDatabase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Init tables if not exists
	tableStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		dodate TEXT,
		finished INTEGER NOT NULL
	);
	`

	_, err = db.Exec(tableStmt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

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

func AddTodo(db *sql.DB, todo *Todo) error {
	stmt := `
	INSERT INTO todos (title, description, dodate, finished)
	VALUES (?, ?, ?, 0);
	`
	_, err := db.Exec(stmt, todo.Title, todo.Description, todo.Dodate)
	if err != nil {
		return err
	}

	return nil
}

func GetTodos(db *sql.DB) ([]*Todo, error) {
	stmt := `
	SELECT * FROM todos;
	`
	
	rows, err := db.Query(stmt)
	if err != nil {
		return make([]*Todo, 0), err
	}
	defer rows.Close()
	
	todos := []*Todo{}
	for rows.Next() {
		todo := &Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Dodate, &todo.Finished)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func FinishTodo(db *sql.DB, id int) error {
	stmt := `
	UPDATE todos
	SET finished = 1
	WHERE id=?;
	`
	_, err := db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func UnfinishTodo(db *sql.DB, id int) error {
	stmt := `
	UPDATE todos
	SET finished = 0
	WHERE id=?;
	`
	_, err := db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}