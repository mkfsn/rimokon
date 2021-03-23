package repository

import (
	"context"

	"github.com/mkfsn/rimokon/internal/domain"
)

type unblockTvRepository struct {
	deviceName string
	serialPort domain.SerialPortRepository
}

func NewUnblockTVRepository(serialPort domain.SerialPortRepository) domain.UnblockTVRepository {
	return &unblockTvRepository{
		deviceName: "UnblockTV",
		serialPort: serialPort,
	}
}

func (s *unblockTvRepository) PressPower(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressPower")
}

func (s *unblockTvRepository) PressInput(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressInput")
}

func (s *unblockTvRepository) PressLeft(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressLeft")
}

func (s *unblockTvRepository) PressRight(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressRight")
}

func (s *unblockTvRepository) PressVolumeUp(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressVolumeUp")
}

func (s *unblockTvRepository) PressVolumeDown(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressVolumeDown")
}

func (s *unblockTvRepository) PressEnter(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressEnter")
}

func (s *unblockTvRepository) PressUp(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressUp")
}

func (s *unblockTvRepository) PressDown(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressDown")
}

func (s *unblockTvRepository) PressOK(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressOK")
}

func (s *unblockTvRepository) PressHome(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressHome")
}

func (s *unblockTvRepository) PressMenu(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressMenu")
}

func (s *unblockTvRepository) PressReturn(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressReturn")
}
