package todoServer

import (
	"database/sql"
	"log"
	"net/http"
)

type ToDoServer struct {
	Addr string
	DB   *sql.DB
	frontEnd string
}

func NewServer(addr string, db *sql.DB, frontEndDir string) *ToDoServer {
	return &ToDoServer{
		Addr: addr,
		DB:   db,
		frontEnd: frontEndDir,
	}
}

func (s *ToDoServer) Run() error {
	router := http.NewServeMux()
	s.SetRoutes(router)

	server := http.Server{
		Addr: s.Addr,
		Handler: router,
	}
	log.Printf("Server started. Listening on %s\n", s.Addr)

	return server.ListenAndServe()
}