package port

import "biathlon-competitions/internal/entity"

type (
	ReaderEvents interface {
		ReadLines() ([]string, error)
	}
	ReaderConfig interface {
		ReadConfig() (*entity.Config, error)
	}
)
