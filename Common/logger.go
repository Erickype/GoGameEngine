package Common

import common "github.com/Erickype/GoGameEngine/Log"

var CoreLogger *common.CoreLogger
var ClientLogger *common.ClientLogger

func initLoggerSystem() {
	CoreLogger = &common.CoreLogger{}
	CoreLogger.Init()

	ClientLogger = &common.ClientLogger{}
	ClientLogger.Init()
}
