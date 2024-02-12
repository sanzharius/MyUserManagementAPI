package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"myAPIProject/internal/apperrors"
)

type Config struct {
	Mongo *Mongo
	Port  string `env:"PORT"`
}

type Mongo struct {
	DBHost       string `env:"MONGODB_URI"`
	DBName       string `env:"DB_NAME"`
	DBPrefix     string `env:"DB_PREFIX"`
	DBUsername   string `env:"DB_USERNAME"`
	DBPassword   string `env:"DB_PASSWORD"`
	DBCollection string `env:"DB_COLLECTION"`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, apperrors.EnvConfigLoadErr.AppendMessage(err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, apperrors.EnvConfigParseErr.AppendMessage(err)
	}

	MongoConfig := &Mongo{}
	if err := env.Parse(MongoConfig); err != nil {
		return nil, apperrors.EnvConfigParseErr.AppendMessage(err)
	}

	cfg.Mongo = MongoConfig

	fmt.Printf("%+v\n", cfg)

	return cfg, nil

}
