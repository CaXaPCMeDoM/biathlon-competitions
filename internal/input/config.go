package input

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/infrastructure"
)

type Config struct {
	reader infrastructure.ReaderConfig
}

func NewConfig(reader infrastructure.ReaderConfig) *Config {
	return &Config{
		reader: reader,
	}
}

func (c *Config) LoadConfig() (*entity.Config, error) {
	return c.reader.ReadConfig()
}
