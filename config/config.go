package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"path"
	"time"
)

type (
	Config struct {
		HTTP  HTTP  `yaml:"http"`
		Log   Log   `yaml:"log"`
		Auth  Auth  `yaml:"auth"`
		Cache Cache `yaml:"cache"`
		PG    PG
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Auth struct {
		SignKey  string        `env:"JWT_SIGN_KEY"`
		TokenTTL time.Duration `yaml:"token_ttl"`
	}

	PG struct {
		URL string `env:"POSTGRES_URL"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	Cache struct {
		ExpiredTTl time.Duration `yaml:"expired_ttl"`
		CleanupTTL time.Duration `yaml:"cleanup_ttl"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	return cfg, nil
}
