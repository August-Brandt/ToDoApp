package tododatabase

import (
	"database/sql"

	. "ToDoServer/datatypes"
)

func Setup(databaseVersion, connection string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../database/ToDoDatabase.db")
	if err != nil {
		return db, err
	}

	// Init tables if not exists
	tableStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id TEXT NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		dodate TEXT,
		finished INTEGER NOT NULL
	);
	`

	_, err = db.Exec(tableStmt)
	if err != nil {
		return db, err
	}
	return db, nil
}

func AddTodo(db *sql.DB, todo *Todo) (string, error) {
	stmt := `
	INSERT INTO todos (id, title, description, dodate, finished)
	VALUES (?, ?, ?, ?, 0);
	`
	_, err := db.Exec(stmt, todo.Id, todo.Title, todo.Description, todo.Dodate)
	if err != nil {
		return "", err
	}

	return todo.Id, nil
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

func GetTodoById(db *sql.DB, id string) (*Todo, error) {
	stmt := `
	SELECT * FROM todos WHERE id=?;
	`

	row := db.QueryRow(stmt, id)
	
	todo := &Todo{}
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Dodate, &todo.Finished)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func RemoveTodoById(db *sql.DB, id string) error {
	stmt := `
	DELETE FROM todos
	WHERE id=?;
	`

	_, err := db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}