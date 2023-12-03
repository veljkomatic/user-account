package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/veljkomatic/user-account/common/environment"
	"github.com/veljkomatic/user-account/common/log"
	"github.com/veljkomatic/user-account/common/ptr"
)

const (
	ConfigFileLocal      = "config.local.yaml"
	ConfigFileProduction = "config.production.yaml"
	ConfigFileCanary     = "config.canary.yaml"
)

type ConfigBuilder[T Validator] struct {
	serviceName string
	shouldInit  bool
	envAccessor environment.EnvAccessor
}

func Builder[T Validator]() *ConfigBuilder[T] {
	envAccessor := environment.RealEnvAccessor{}
	return &ConfigBuilder[T]{
		serviceName: environment.ServiceName(ptr.From(envAccessor)),
		shouldInit:  true,
		envAccessor: ptr.From(envAccessor),
	}
}

func (cb *ConfigBuilder[T]) Build(ctx context.Context) T {
	configFilePath := cb.getConfigFilePath()

	cfg, err := cb.LoadConfig(configFilePath)
	if err != nil {
		log.Error(ctx, err, "failed to load config")
		panic(err)
	}

	if err := cfg.Validate(); err != nil {
		log.Error(ctx, err, "failed to validate config")
		panic(err)
	}

	log.Info(
		ctx,
		"config loaded",
		log.Fields{
			"service_name":    environment.ServiceName(cb.envAccessor),
			"service_version": environment.ServiceVersion(cb.envAccessor),
		},
	)
	return cfg
}

func (cb *ConfigBuilder[T]) LoadConfig(path string) (T, error) {
	var cfg T

	data, err := os.ReadFile(path)
	if err != nil {
		return ptr.Value(new(T)), err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return ptr.Value(new(T)), err
	}

	return cfg, nil
}

// getConfigFilePath determines the configuration file path based on the environment
func (cb *ConfigBuilder[T]) getConfigFilePath() string {
	configDir := cb.envAccessor.GetEnv("USER_ACCOUNT_CONFIG_PATH")
	if configDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Warn(context.TODO(), fmt.Sprintf("Error getting current working directory: %s", err.Error()))
			// Handle the error, maybe fallback to a default config directory
		}
		configDir = cwd // Use the current working directory
	}
	currentEnv := environment.Get(cb.envAccessor)

	var configFile string
	switch currentEnv {
	case environment.EnvProduction:
		configFile = ConfigFileProduction
	case environment.EnvCanary:
		configFile = ConfigFileCanary
	default:
		configFile = ConfigFileLocal
	}

	return filepath.Join(configDir, configFile)
}
