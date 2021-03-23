package repository

import (
	"context"

	"github.com/mkfsn/rimokon/internal/domain"
)

type logiSpeakerRepository struct {
	deviceName string
	serialPort domain.SerialPortRepository
}

func NewLogiSpeakerRepository(serialPort domain.SerialPortRepository) domain.LogiSpeakerRepository {
	return &logiSpeakerRepository{
		deviceName: "LogiSpeaker",
		serialPort: serialPort,
	}
}

func (s *logiSpeakerRepository) PressPower(ctx context.Context) error {
	return s.serialPort.Execute(ctx, s.deviceName, "pressPower")
}
