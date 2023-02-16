package Core

import (
	"github.com/Erickype/GoGameEngine/Log"
	"time"
)

type IApplication interface {
	run()
	destroy()
	init()
}

type Application struct {
	coreLogger   *Log.CoreLogger
	clientLogger *Log.ClientLogger
}

func (a *Application) run() {
	for {
		time.Sleep(time.Second)
		a.clientLogger.Info("Running")
	}
}

func (a *Application) destroy() {
}

func (a *Application) init() {
	a.coreLogger, a.clientLogger = InitLogSystem()
	a.coreLogger.Info("Starting engine!!")
}

// CreateApplication This is the entry point to create an application
func CreateApplication() {
	application := &Application{}
	application.init()
	application.run()
	application.destroy()
}
