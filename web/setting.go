package web

import "net/http"


func (s *Server) SettingHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/setting/index", "default", nil)
}
