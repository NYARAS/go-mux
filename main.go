package main

import (
	"github.com/NYARAS/go-mux/app"
	"github.com/NYARAS/go-mux/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
