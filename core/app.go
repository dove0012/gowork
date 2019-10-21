package core

import (
	"gowork/common/utils/log"
	mtime "gowork/common/utils/time"
	"gowork/work"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type App struct {
	Name    string
	Usage   string
	Version string
	Works   []work.Work
	Wg      sync.WaitGroup
}

func NewApp() *App {
	return &App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "A new cli application",
		Version: "0.0.0",
	}
}

func (app *App) Run() {
	log.Info("App starting")
	runtime.GOMAXPROCS(runtime.NumCPU())
	for _, server := range app.Works {
		app.Wg.Add(1)
		app.runWork(server)
	}
	app.Wg.Wait()
}

func (app *App) rebootWork(name string) {
	log.Info("work[" + name + "] rebooting")
	for _, work := range app.Works {
		if name == work.GetName() {
			app.runWork(work)
		}
	}
}

func (app *App) runWork(Work work.Work) {
	log.Info("work[" + Work.GetName() + "] running")
	go func() {
		startTime := mtime.NowUnixMilli()
		if Work.GetReboot() {
			defer func() {
				log.TimeConsuming(startTime, "["+Work.GetName()+"] is over")
				time.Sleep(time.Second * Work.GetRebootTime())
				app.rebootWork(Work.GetName())
			}()
		} else {
			defer func() {
				log.TimeConsuming(startTime, "["+Work.GetName()+"] is over")
				app.Wg.Done()
			}()
		}
		Work.Run()
	}()
}
