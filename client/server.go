package client

import (
	"fmt"
	"govie.io/govie-server/core"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) Init() {
	s.Router = http.NewServeMux()
	s.Router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./client-web/assets"))))
	s.Router.HandleFunc("GET /login", s.LoginHandler)
	s.Router.HandleFunc("GET /{$}", core.AuthMiddleware(s.Handle, "/login"))

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: s.Router,
	}

	// launch server
	fmt.Println("Web Server : Start : ", server.ListenAndServe())

}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HOME : WEB SERVER!<br /><a href='/login'>Login</a>")
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "LOGIN : WEB SERVER!")
}
