package input

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/port"
)

type Config struct {
	reader port.ReaderConfig
}

func NewConfig(reader port.ReaderConfig) *Config {
	return &Config{
		reader: reader,
	}
}

func (c *Config) LoadConfig() (*entity.Config, error) {
	return c.reader.ReadConfig()
}
