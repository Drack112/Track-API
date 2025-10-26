package main

import (
	"os"

	"github.com/Drack112/Track-API/internal/adapter/secondary/contextlogger"
)

const (
	MsgConfigLoaded    = "configuration loaded"
	MsgDepsInitialized = "dependencies initialized"
	ErrLoadConfig      = "failed to load configuration"
	ErrInitDeps        = "failed to initialize dependencies"
	ErrServerRunFailed = "server run failed"
	SwaggerTitle       = "TrackAPI - REST API"
)

func main() {
	os.Exit(run())
}

func run() int {
	logs, cleanupLogger := contextlogger.New()
	defer cleanupLogger()

	logs.Infof("Running")

	return 1
}
