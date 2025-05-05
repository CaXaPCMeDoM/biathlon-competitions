package json

import (
	"biathlon-competitions/internal/entity"
	"encoding/json"
	"os"
)

type ConfigReader struct {
	filepath string
}

func New(filepath string) *ConfigReader {
	return &ConfigReader{
		filepath: filepath,
	}
}

func (r *ConfigReader) ReadConfig() (*entity.Config, error) {
	data, err := os.ReadFile(r.filepath)
	if err != nil {
		return nil, err
	}

	var cfg entity.Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
