package main

import (
	"gowork/common/utils/log"
	"gowork/core"
	"gowork/work"
)

const AppVer = "1.0.0"

func main() {
	app := core.NewApp()
	app.Name = "gowork"
	app.Usage = "run services by go"
	app.Version = AppVer
	app.Works = work.GetWorks()
	app.Run()
	log.Info("Application is over")
}
