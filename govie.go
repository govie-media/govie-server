package main

import (
	"embed"
	"govie.io/govie-server/api"
	"govie.io/govie-server/core"
	"govie.io/govie-server/image"
	"govie.io/govie-server/web"
	"io/fs"
)

type Govie struct {
	Version string

	// Servers
	WebServer   *web.Server
	ApiServer   *api.Server
	ImageServer *image.Server

	// Settings
	Settings core.GovieSettings

	Tasks string
}

//go:embed webroot/assets/* webroot/layout/* webroot/view/*
var staticWebFiles embed.FS

func (g *Govie) Init(disableWebServer, disableApiServer, disableImageServer bool) {
	g.Version = "0.0.1"

	// Read settings

	// Start Web Server
	if !disableWebServer {
		go func() {
			// Get assets within the webroot
			files, _ := fs.Sub(staticWebFiles, "webroot")

			// Start Server
			g.WebServer = &web.Server{}
			g.WebServer.Init(files)
		}()
	}

	// Start Api Server
	if !disableApiServer {
		go func() {
			g.ApiServer = &api.Server{}
			g.ApiServer.Init()
		}()
	}

	// Start Image Server
	if !disableImageServer {
		go func() {
			g.ImageServer = &image.Server{}
			g.ImageServer.Init()
		}()
	}

	// Wait
	select {}
}
