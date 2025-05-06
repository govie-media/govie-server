package web

import (
	"errors"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"govie.io/govie-server/core"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	HTTPServer *http.Server
	Router     *http.ServeMux
	Settings   *core.WebSettings
	Validator  *validator.Validate
	Decoder    *form.Decoder
}

func (s *Server) Init(staticFiles fs.FS) {
	loginUrl := "/login"

	s.Settings.Webroot = "./webroot"
	s.Validator = validator.New()
	s.Decoder = form.NewDecoder()

	s.Router = http.NewServeMux()
	// TODO: SWITCH TO EMBEDDED FILESYSTEM
	//s.Router.Handle("/assets/", core.NeuterHttpFileServer(http.FileServerFS(staticFiles)))
	s.Router.Handle("/assets/", http.FileServer(http.Dir(s.Settings.Webroot)))

	// Setup
	s.Router.HandleFunc("/setup", s.SetupHandler)

	// Account
	s.Router.HandleFunc(loginUrl, s.LoginHandler)
	s.Router.HandleFunc("/logout", s.LogoutHandler)
	s.Router.HandleFunc("GET /account", core.ValidateCookieAuth(s.AccountHandler, loginUrl))

	// Library
	s.Router.HandleFunc("GET /library", core.ValidateCookieAuth(s.LibraryHandler, loginUrl))

	// Users
	s.Router.HandleFunc("GET /user", core.ValidateCookieAuth(s.UsersHandler, loginUrl))
	s.Router.HandleFunc("/user/{test}", core.ValidateCookieAuth(s.EditUserHandler, loginUrl))

	// Settings
	s.Router.HandleFunc("GET /plugin", core.ValidateCookieAuth(s.PluginHandler, loginUrl))
	s.Router.HandleFunc("GET /log", core.ValidateCookieAuth(s.LogHandler, loginUrl))
	s.Router.HandleFunc("GET /setting", core.ValidateCookieAuth(s.SettingHandler, loginUrl))

	// Dashboard
	s.Router.HandleFunc("GET /search", core.ValidateCookieAuth(s.SearchHandler, loginUrl))
	s.Router.HandleFunc("GET /{$}", core.ValidateCookieAuth(s.DashboardHandler, loginUrl))

	// create server to run on port the 9000
	s.HTTPServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.Settings.Port),
		Handler: s.Router,
	}

	// launch server
	fmt.Println("Web Server : Start : ", s.HTTPServer.ListenAndServe())
}

func (s *Server) Render(w http.ResponseWriter, view, layout string, data interface{}) {
	viewTmpl := s.Settings.Webroot + "/view" + view + ".html"
	layoutTmpl := s.Settings.Webroot + "/layout/" + layout + ".html"

	t, err := template.ParseFiles(layoutTmpl, viewTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) ValidateForm(r *http.Request, target interface{}) error {
	// Ensure target is a pointer to a struct
	//rv := reflect.ValueOf(target)
	//if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
	//	return errors.New("target must be a pointer to a struct")
	//}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		return errors.New("failed to parse form")
	}

	// Decode form data into struct
	err := s.Decoder.Decode(target, r.Form)
	if err != nil {
		return errors.New("form decode failed")
	}

	// Validate the form
	err = s.Validator.Struct(target)
	if err != nil {
		log.Println(err.Error())
		return errors.New("validation failed")
	}

	return nil
}

func (s *Server) Shutdown() {

}
