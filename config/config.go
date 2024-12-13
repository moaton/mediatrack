package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap/zapcore"
)

type (
	Config struct {
		Logger   `yaml:"logger"`
		Postgres `yaml:"postgres"`
		HTTP     `yaml:"http"`
	}

	Logger struct {
		Level zapcore.Level `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Postgres struct {
		Host     string `env-required:"true" yaml:"postgres_host" env:"DB_HOST"`
		Port     uint16 `env-required:"true" yaml:"postgres_port" env:"DB_PORT"`
		User     string `env-required:"true" yaml:"postgres_user" env:"DB_USER"`
		Password string `env-required:"true" yaml:"postgres_password" env:"DB_PASSWORD"`
		DBName   string `env-required:"true" yaml:"postgres_db_name" env:"DB_NAME"`
	}

	HTTP struct {
		Server struct {
			Port string `env-required:"true" yaml:"http_server_port" env:"HTTP_SERVER_PORT"`
		} `yaml:"server"`
	}
)

func (p Postgres) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", url.QueryEscape(p.User), url.QueryEscape(p.Password), p.Host, p.DBName)
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to load env from file: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
