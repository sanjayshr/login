package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server is type struct for the server

type Server struct {
	Router *mux.Router
}

// Initialize will initialize DB and Router
func (server *Server) Initialize(DBName string) {

	fmt.Println("DB Name: %s", DBName)

	server.Router = mux.NewRouter()
	server.initializeRoutes()

}

// Run will spin the server and pass routes
func (server *Server) Run(addr string) {

	fmt.Println("Listening on port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
