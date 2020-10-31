package server

import "net/http"

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

}