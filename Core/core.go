package Core

import "github.com/Erickype/GoGameEngine/Log"

func InitLogSystem() (*Log.CoreLogger, *Log.ClientLogger) {

	coreLogger := &Log.CoreLogger{}
	coreLogger.Init()

	clientLogger := &Log.ClientLogger{}
	clientLogger.Init()

	return coreLogger, clientLogger
}
