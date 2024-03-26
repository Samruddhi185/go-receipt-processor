package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	router.HandleFunc("/receipts/process", makeHTTPHandlerFunc(s.handleReceiptProcess)).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", makeHTTPHandlerFunc(s.handleReceiptGetPoints)).Methods("GET")

	log.Println("Receipt server running on port: ", s.listenPort)
	http.ListenAndServe(s.listenPort, router)
}

func (s *Server) handleReceiptProcess(w http.ResponseWriter, r *http.Request) error {
	if r.Body == nil {
		log.Println("request body is empty")
		return WriteJSON(w, http.StatusBadRequest, "The receipt is invalid")
	}

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		log.Println("Unable to read JSON. Encountered error: ", err)
		return WriteJSON(w, http.StatusBadRequest, "The receipt is invalid")
	}

	return WriteJSON(w, http.StatusOK, err)
}

func (s *Server) handleReceiptGetPoints(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func makeHTTPHandlerFunc(h apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
