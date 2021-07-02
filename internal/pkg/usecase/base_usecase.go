package usecase

import "go-api/pkg/infrastructure"

// BaseUsecase struct.
type BaseUsecase struct {
	Logger infrastructure.Logger
}

// NewBaseUsecase returns NewBaseUsecase instance.
func NewBaseUsecase(logger infrastructure.Logger) *BaseUsecase {
	return &BaseUsecase{Logger: logger}
}
