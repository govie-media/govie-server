package main

import "flag"

func main() {
	// Command Arguments
	disableWebServer := flag.Bool("disableWebServer", false, "(Optional, Boolean) Flag to disable the govie web server")
	disableApiServer := flag.Bool("disableApiServer", false, "(Optional, Boolean) Flag to disable the govie api server")
	disableImageServer := flag.Bool("disableImageServer", false, "(Optional, Boolean) Flag to disable the govie image server")
	flag.Parse()

	// Start Govie Servers
	govie := &Govie{}

	govie.Init(
		*disableWebServer,
		*disableApiServer,
		*disableImageServer,
	)
}
