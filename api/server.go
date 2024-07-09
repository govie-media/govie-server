package api

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s *Server) Init() {
	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9001",
		Handler: http.HandlerFunc(s.Handle),
	}

	// launch server
	fmt.Println("ListenAndServe():", server.ListenAndServe())
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API SERVER!")
}
