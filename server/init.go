package server

import (
	"github.com/PingTeeratorn789/reverse_proxy/configs"
	logger "github.com/PingTeeratorn789/reverse_proxy/internal/pkg/loggerr"
)

var (
	app *application
)

type application struct {
	confs  *configs.Config
	logger *logger.Logger
}

func init() {
	config := configs.Init()
	logger := logger.NewLogger(logger.Dependencies{Configs: config})
	app = &application{
		confs:  config,
		logger: logger,
	}
}
