package server

import (
	"github.com/gorilla/mux"
	"github.com/hollowsunsets/Go-REST-API/config"
)

type Server struct {
	config *config.Config
	router *mux.Router
}

func New(c *config.Config) (*Server, error) {
	s := &Server{
		config: c,
		router: mux.NewRouter(),
	}

	s.router.HandleFunc("/login", s.Login).Methods("POST")
	s.router.HandleFunc("/signup", s.SignUp).Methods("POST")

	return s, nil
}


