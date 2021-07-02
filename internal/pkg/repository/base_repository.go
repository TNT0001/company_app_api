package repository

import "go-api/pkg/infrastructure"

// BaseRepository struct.
type BaseRepository struct {
	Logger infrastructure.Logger
}

// NewBaseRepository returns NewBaseRepository instance.
func NewBaseRepository(logger infrastructure.Logger) *BaseRepository {
	return &BaseRepository{Logger: logger}
}
