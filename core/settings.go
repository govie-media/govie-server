package core

type GovieSettings struct {
	Api   Settings
	Image Settings
	Web   Settings
}

type Settings struct {
	Port    string
	Webroot string
}
