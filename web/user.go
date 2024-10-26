package web

import "net/http"

func (s *Server) UsersHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/user/index", "default", nil)
}
