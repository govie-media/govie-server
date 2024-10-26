package web

import "net/http"

func (s *Server) LogHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/log/index", "default", nil)
}
