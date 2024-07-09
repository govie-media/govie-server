package web

import (
	"fmt"
	"govie.io/govie-server/core"
	"html/template"
	"io/fs"
	"net/http"
)

type Server struct {
	HTTPServer *http.Server
	Router     *http.ServeMux
	Settings   *core.Settings
}

func (s *Server) Init(staticFiles fs.FS) {
	s.Settings = &core.Settings{}
	s.Settings.Webroot = "./webroot"

	s.Router = http.NewServeMux()
	s.Router.Handle("/assets/", core.NeuteredFileSystemIntercept(http.FileServerFS(staticFiles)))
	s.Router.HandleFunc("/login", s.LoginHandler)
	s.Router.HandleFunc("GET /search", core.HttpAuth(s.SearchHandler, "/login"))
	s.Router.HandleFunc("GET /{$}", core.HttpAuth(s.HomeHandler, "/login"))

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

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	s.Render(w, "/main/index", "default", nil)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		token, err := core.ParseAuthRequest(r)

		if err != nil {
			println(err.Error())
		}

		println(token)
	}

	s.Render(w, "/account/login", "gateway", nil)
}

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Search : WEB SERVER!")
}
