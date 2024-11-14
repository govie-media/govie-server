package web

import "net/http"

func (s *Server) SetupHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/setup/index", "basic", nil)
}
