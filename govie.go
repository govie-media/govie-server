package main

import (
	"govie.io/govie-server/api"
	web "govie.io/govie-server/client"
	"govie.io/govie-server/image"
	"govie.io/govie-server/web"
)

type Govie struct {
	Version string

	// Servers
	WebServer   *web.Server
	ApiServer   *api.Server
	ImageServer *image.Server

	Tasks string
}

func (g *Govie) Init(disableWebServer, disableApiServer, disableImageServer bool) {
	g.Version = "0.0.1"

	// Start Web Server
	if !disableWebServer {
		go func() {
			g.WebServer = &web.Server{}
			g.WebServer.Init()
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
