package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"

	"govie.io/govie-server/api"
	"govie.io/govie-server/core"
	"govie.io/govie-server/image"
	"govie.io/govie-server/web"
)

type Govie struct {
	Version string

	// Servers
	WebServer   *web.Server
	ApiServer   *api.Server
	ImageServer *image.Server

	// Settings
	Settings core.GovieSettings

	// Sync
	Sync *core.Sync

	Tasks string
}

//go:embed webroot/assets/* webroot/layout/* webroot/view/*
var staticWebFiles embed.FS

func (g *Govie) Init(disableWebServer, disableApiServer, disableImageServer bool) {
	g.Version = "2.0.0"

	fmt.Printf("Starting Govie Webservers v%s\n", g.Version)

	// Read settings
	g.LoadSettings()

	// Start Web Server
	if !disableWebServer {
		go func() {
			// Get assets within the webroot
			files, _ := fs.Sub(staticWebFiles, "webroot")

			// Start Server
			g.WebServer = &web.Server{}
			g.WebServer.Settings = &g.Settings.Web
			g.WebServer.Init(files)
		}()
	}

	// Start Api Server
	if !disableApiServer {
		go func() {
			g.ApiServer = &api.Server{}

			g.ApiServer.Settings = &g.Settings.Api
			g.ApiServer.Init()
		}()
	}

	// Start Image Server
	if !disableImageServer {
		go func() {
			g.ImageServer = &image.Server{}
			g.ImageServer.Settings = &g.Settings.Image
			g.ImageServer.Init()
		}()
	}

	// Setup Scheduler
	g.Sync = &core.Sync{}

	// Wait
	if !disableWebServer || !disableApiServer || !disableImageServer {
		// Block until user sends SIGINT or SIGTERM
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs

		g.Shutdown()
	}

	fmt.Println("END")
}

func (g *Govie) LoadSettings() {
	// Open the JSON file
	file, err := os.Open("settings/config.json")
	if err != nil {
		log.Fatalln("Error opening settings file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after reading

	// Decode the JSON data into the struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&g.Settings)
	if err != nil {
		log.Fatalln("Failed to decode settings file:", err)
		return
	}
}

// TODO: Shutdown Servers
func (g *Govie) Shutdown() {

}
