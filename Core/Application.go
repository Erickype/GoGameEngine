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
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		a.clientLogger.Info("Running")
	}
}

func (a *Application) destroy() {
	a.coreLogger.Fatal("Destroying: ", true)
}

func (a *Application) init() {
	a.coreLogger, a.clientLogger = InitLogSystem()
	a.coreLogger.Info("Starting engine!!")
	a.coreLogger.Warn("Develop version: ", true)
	a.coreLogger.Error("Errors: ", false)
	a.coreLogger.Debug("Debug: ", true)
	a.coreLogger.Trace("Trace")
}

// CreateApplication This is the entry point to create an application
func CreateApplication() {
	application := &Application{}
	application.init()
	application.run()
	application.destroy()
}
