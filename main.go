package main

import (
	"github.com/NYARAS/go-mux/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
