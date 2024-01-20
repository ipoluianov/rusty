package main

import (
	"github.com/ipoluianov/rusty/app"
	"github.com/ipoluianov/rusty/application"
	"github.com/ipoluianov/rusty/logger"
)

func main() {
	application.Name = "rusty"
	application.ServiceName = "rusty"
	application.ServiceDisplayName = "rusty"
	application.ServiceDescription = "rusty"
	application.ServiceRunFunc = app.RunAsService
	application.ServiceStopFunc = app.StopService

	logger.Init(logger.CurrentExePath() + "/logs")

	if !application.TryService() {
		app.RunDesktop()
	}
}
