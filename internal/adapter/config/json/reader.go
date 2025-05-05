package json

import (
	"biathlon-competitions/internal/entity"
	"encoding/json"
	"os"
)

type Reader struct {
	filepath string
}

func New(filepath string) *Reader {
	return &Reader{
		filepath: filepath,
	}
}

func (r *Reader) ReadConfig() (*entity.Config, error) {
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
