package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store //repository having the connections to the database
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() { //initialize the router then register services
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//register service
	tasksService := newTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting the API Server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
