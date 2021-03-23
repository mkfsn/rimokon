package repository

import (
	"context"

	"github.com/mkfsn/rimokon/internal/domain"
)

type sharpTvRepository struct {
	deviceName string
	serialPort domain.SerialPortRepository
}

func NewSharpTVRepository(serialPort domain.SerialPortRepository) domain.SharpTVRepository {
	return &sharpTvRepository{
		deviceName: "SharpTV",
		serialPort: serialPort,
	}
}

func (s *sharpTvRepository) PressPower(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressPower")
}

func (s *sharpTvRepository) PressInput(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressInput")
}

func (s *sharpTvRepository) PressLeft(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressLeft")
}

func (s *sharpTvRepository) PressRight(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressRight")
}

func (s *sharpTvRepository) PressVolumeUp(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressVolumeUp")
}

func (s *sharpTvRepository) PressVolumeDown(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressVolumeDown")
}

func (s *sharpTvRepository) PressEnter(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressEnter")
}
