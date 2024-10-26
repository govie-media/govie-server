package web

import "net/http"

func (s *Server) PluginHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/plugin/index", "default", nil)
}
