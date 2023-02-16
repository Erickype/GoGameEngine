package Log

import "github.com/withmandala/go-log"

type ClientLogger struct {
	logger *log.Logger
}

func (c *ClientLogger) Init() {
	c.logger = InitLogger()
}

func (c *ClientLogger) Info(message string) {
	c.logger.Info(message)
}
