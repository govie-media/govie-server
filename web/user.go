package web

import (
	"fmt"
	"net/http"
)

type NewUserForm struct {
	Username string `form:"username" validate:"required,min=3,max=254"`
	Alias    string `form:"alias" validate:"max=50"`
	Password string `form:"password" validate:"required"`
}

func (s *Server) UsersHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/user/index", "default", nil)
}

func (s *Server) EditUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var form NewUserForm

		err := s.ValidateForm(r, &form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create User
		fmt.Printf("New User: %v\n", form)

		// Redirect
		http.Redirect(w, r, "/user", http.StatusFound)
	}

	s.Render(w, "/user/edit", "default", nil)
}
