package core

type GovieSettings struct {
	Api      WebSettings `json:"api"`
	Image    WebSettings `json:"image"`
	Web      WebSettings `json:"web"`
	Database DatabaseSettings
}

type WebSettings struct {
	Port    int
	Webroot string
}

type DatabaseSettings struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Type     string
}
