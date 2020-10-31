package server

import "net/http"

type signupRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email	  string `json:"email"`
	Password  string `json:"password"`
}

type signupResponse struct {
	Token string `json:"token"`
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {

}