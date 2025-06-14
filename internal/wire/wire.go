//go:build wireinject

package wire

import (
	"github.com/PRSV-Hackathon-2025/go-backend/internal/controller"
	"github.com/google/wire"
)

func InitializeHelloController() (*controller.HelloController, error) {
	wire.Build(
		controller.NewHelloController,
	)

	return &controller.HelloController{}, nil
}
