package main

import (
	"github.com/gawelczyk/api1/app"
	"github.com/gawelczyk/api1/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
