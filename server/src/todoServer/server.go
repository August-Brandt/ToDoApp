package todoServer

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"
)

type ToDoServer struct {
	Addr     string
	DB       *sql.DB
	frontEnd string
	Mode     string
}

func NewServer(addr string, db *sql.DB, frontEndDir string) *ToDoServer {
	mode, exists := os.LookupEnv("SERVERMODE")
	if !exists {
		mode = "PROD"
	}

	return &ToDoServer{
		Addr:     addr,
		DB:       db,
		frontEnd: frontEndDir,
		Mode:     strings.ToUpper(mode),
	}
}

func (s *ToDoServer) Run() error {
	router := http.NewServeMux()
	s.SetRoutes(router)

	server := http.Server{
		Addr:    s.Addr,
		Handler: router,
	}
	log.Printf("Server started in %s mode. Listening on %s\n", s.Mode, s.Addr)

	return server.ListenAndServe()
}
