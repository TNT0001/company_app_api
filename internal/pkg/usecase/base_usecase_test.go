package usecase

import (
	"go-api/pkg/shared/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseUseCase(t *testing.T) {
	logger := new(test.LoggerMock)
	bu := NewBaseUsecase(logger)
	assert.IsType(t, &BaseUsecase{}, bu)
}
