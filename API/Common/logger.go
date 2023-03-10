package Common

import (
	"github.com/Erickype/GoGameEngine/API/Log"
)

var CoreLogger *Log.CoreLogger
var ClientLogger *Log.ClientLogger

func initLoggerSystem() {
	CoreLogger = &Log.CoreLogger{}
	CoreLogger.Init()

	ClientLogger = &Log.ClientLogger{}
	ClientLogger.Init()
}
