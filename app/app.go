package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	listenPort string
}

type apiHandler func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func CreateServer(listenPort string) *Server {
	return &Server{
		listenPort: listenPort,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	log.Println("Receipt server running on port: ", s.listenPort)
	http.ListenAndServe(s.listenPort, router)
}
