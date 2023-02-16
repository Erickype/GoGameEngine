package Log

import "github.com/withmandala/go-log"

type CoreLogger struct {
	logger *log.Logger
}

func (t *CoreLogger) Init() {
	t.logger = InitLogger()
}

func (t *CoreLogger) Info(message string) {
	t.logger.Info(message)
}
