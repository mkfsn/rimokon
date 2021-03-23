package domain

import (
	"context"
)

type RimokonUsecase interface {
	ExecuteDeviceCommand(ctx context.Context, string, command string) error
}
