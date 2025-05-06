package api

import (
	"fmt"
	"govie.io/govie-server/core"
	"net/http"
	"strconv"
)

type Server struct {
	Settings *core.WebSettings
}

func (s *Server) Init() {
	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(s.Settings.Port),
		Handler: http.HandlerFunc(s.Handle),
	}

	// launch server
	fmt.Println("ListenAndServe():", server.ListenAndServe())
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API SERVER!")
}
