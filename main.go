package main

import (
	"core"
	"work"
	"common/utils/log"
)

const APP_VER = "1.0.0"

func main() {
	app := core.NewApp()
	app.Name = "goserver"
	app.Usage = "run services by go"
	app.Version = APP_VER
	app.Works = work.GetWorks()
	app.Run()
	log.Info("App is over")
}