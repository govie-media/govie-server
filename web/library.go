package web

import "net/http"

func (s *Server) LibraryHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/library/index", "default", nil)
}
