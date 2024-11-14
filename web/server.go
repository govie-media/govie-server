package web

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"govie.io/govie-server/core"
)

type Server struct {
	HTTPServer *http.Server
	Router     *http.ServeMux
	Settings   *core.Settings
}

func (s *Server) Init(staticFiles fs.FS) {
	s.Settings = &core.Settings{}
	s.Settings.Webroot = "./webroot"
	loginUrl := "/login"

	s.Router = http.NewServeMux()
	// TODO: SWITCH TO EMBEDDED FILESYSTEM
	//s.Router.Handle("/assets/", core.NeuterHttpFileServer(http.FileServerFS(staticFiles)))
	s.Router.Handle("/assets/", http.FileServer(http.Dir(s.Settings.Webroot)))

	// Setup
	s.Router.HandleFunc("/setup", s.SetupHandler)

	// Account
	s.Router.HandleFunc(loginUrl, s.LoginHandler)
	s.Router.HandleFunc("/logout", s.LogoutHandler)
	s.Router.HandleFunc("/account", core.ValidateCookieAuth(s.AccountHandler, loginUrl))

	// Library
	s.Router.HandleFunc("/library", core.ValidateCookieAuth(s.LibraryHandler, loginUrl))

	// Users
	s.Router.HandleFunc("/user", core.ValidateCookieAuth(s.UsersHandler, loginUrl))

	// Settings
	s.Router.HandleFunc("/plugin", core.ValidateCookieAuth(s.PluginHandler, loginUrl))
	s.Router.HandleFunc("/log", core.ValidateCookieAuth(s.LogHandler, loginUrl))
	s.Router.HandleFunc("/setting", core.ValidateCookieAuth(s.SettingHandler, loginUrl))

	// Dashboard
	s.Router.HandleFunc("GET /search", core.ValidateCookieAuth(s.SearchHandler, loginUrl))
	s.Router.HandleFunc("GET /{$}", core.ValidateCookieAuth(s.DashboardHandler, loginUrl))

	// create server to run on port the 9000
	s.HTTPServer = &http.Server{
		Addr:    ":9000",
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
