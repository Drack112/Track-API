package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Drack112/Track-API/internal/adapter/secondary/contextlogger"
	"github.com/Drack112/Track-API/internal/adapter/secondary/crypto"
	"github.com/Drack112/Track-API/internal/platform/config"
	"github.com/Drack112/Track-API/internal/platform/ports/output/keygen"
	"github.com/Drack112/Track-API/internal/platform/ports/output/logger"
	"github.com/Drack112/Track-API/internal/shared/constants/commonkeys"
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

	keyGenerator := crypto.New()

	_, err := loadConfig(keyGenerator, logs)
	if err != nil {
		logs.Errorw(ErrLoadConfig, commonkeys.Error, err.Error())
		return 2
	}

	_, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	return 0
}

func loadConfig(keyGenerator keygen.Generator, logs logger.ContextLogger) (*config.Config, error) {
	cfg, err := config.New(keyGenerator).Load(logs)
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	logs.Infow(
		MsgConfigLoaded,
		commonkeys.APIName, cfg.General.Name,
		commonkeys.AppEnv, cfg.General.Env,
		commonkeys.AppVersion, cfg.General.Version,
	)

	return cfg, nil
}
