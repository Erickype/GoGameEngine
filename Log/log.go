package Log

import (
	"github.com/withmandala/go-log"
	"os"
)

type ILogger interface {
	Init()
	Info(message string)
}

func InitLogger() *log.Logger {
	return log.New(os.Stderr)
}
