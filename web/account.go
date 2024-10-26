package web

import (
	"net/http"
	"time"

	"govie.io/govie-server/core"
)

func (s *Server) AccountHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/account/index", "default", nil)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token, err := core.ParseAuthRequest(r)
		if err != nil {
			println(err.Error())
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "gid",
				Value:    *token,
				Expires:  time.Now().Add(2 * time.Hour),
				HttpOnly: true,
			})
			// Redirect to a different URL
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	} else {
		// Clear the token cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "gid",
			Value:    "",
			Expires:  time.Now(),
			HttpOnly: true,
		})
	}

	s.Render(w, "/account/login", "gateway", nil)
}

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "gid",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
