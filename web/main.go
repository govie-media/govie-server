package web

import (
	"net/http"
)

func (s *Server) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/main/index", "default", nil)
}

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/main/search", "default", nil)
}
