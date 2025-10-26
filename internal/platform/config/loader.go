package config

import (
	"strings"

	"github.com/Drack112/Track-API/internal/platform/ports/output/keygen"
	"github.com/Drack112/Track-API/internal/platform/ports/output/logger"
	"github.com/Drack112/Track-API/internal/shared/constants/commonkeys"
	"github.com/kelseyhightower/envconfig"
)

type Loader struct {
	keyGenerator keygen.Generator
	cfg          Config
}

func New(keyGen keygen.Generator) *Loader {
	return &Loader{
		keyGenerator: keyGen,
	}
}

func (l *Loader) Load(logger logger.ContextLogger) (*Config, error) {
	if err := envconfig.Process(commonkeys.Setting, &l.cfg); err != nil {
		logger.Errorw(ErrFailedToProcessEnvVars, commonkeys.Error, err)
		return nil, err
	}

	if !strings.HasPrefix(l.cfg.ServerHTTP.Context, "/") {
		l.cfg.ServerHTTP.Context = "/" + l.cfg.ServerHTTP.Context
	}

	if l.cfg.Secret.Key == "" {
		generated, err := l.keyGenerator.Generate()
		if err != nil {
			logger.Errorw(ErrGenerateSecretKey, commonkeys.Error, err)
			return nil, err
		}

		l.cfg.Secret.Key = generated

		logger.Warnf(SecretKeyWasNotSet)
		logger.Infof(InfoSecretKeyGenerated, len(generated))
	}

	return &l.cfg, nil
}
