package usecase

import (
	"context"

	"github.com/mkfsn/rimokon/internal/domain"
)

type rimokonUsecase struct {
	sharpTVRepository     domain.SharpTVRepository
	logiSpeakerRepository domain.LogiSpeakerRepository
	unblockTVRepository   domain.UnblockTVRepository
}

func NewRimokonUsecase(
	sharpTVRepository domain.SharpTVRepository,
	logiSpeakerRepository domain.LogiSpeakerRepository,
	unblockTVRepository domain.UnblockTVRepository,
) domain.RimokonUsecase {
	return &rimokonUsecase{
		sharpTVRepository:     sharpTVRepository,
		logiSpeakerRepository: logiSpeakerRepository,
		unblockTVRepository:   unblockTVRepository,
	}
}

func (r *rimokonUsecase) ExecuteDeviceCommand(ctx context.Context, device string, command string) error {
	switch device {
	case "SharpTV":
		return r.executeSharpTVCommand(ctx, command)
	case "LogiSpeaker":
		return r.executeLogiSpeakerCommand(ctx, command)
	case "UnblockTV":
		return r.executeUnblockTVCommand(ctx, command)
	}
	return domain.ErrUnknownDevice
}

func (r *rimokonUsecase) executeSharpTVCommand(ctx context.Context, command string) error {
	switch command {
	case "pressPower":
		return r.sharpTVRepository.PressPower(ctx)
	case "pressInput":
		return r.sharpTVRepository.PressInput(ctx)
	case "pressLeft":
		return r.sharpTVRepository.PressLeft(ctx)
	case "pressRight":
		return r.sharpTVRepository.PressRight(ctx)
	case "pressVolumeUp":
		return r.sharpTVRepository.PressVolumeUp(ctx)
	case "pressVolumeDown":
		return r.sharpTVRepository.PressVolumeDown(ctx)
	case "pressEnter":
		return r.sharpTVRepository.PressEnter(ctx)
	}
	return domain.ErrUnknownCommand
}

func (r *rimokonUsecase) executeLogiSpeakerCommand(ctx context.Context, command string) error {
	switch command {
	case "pressPower":
		return r.logiSpeakerRepository.PressPower(ctx)
	}
	return domain.ErrUnknownCommand
}

func (r *rimokonUsecase) executeUnblockTVCommand(ctx context.Context, command string) error {
	switch command {
	case "pressPower":
		return r.unblockTVRepository.PressPower(ctx)
	case "pressLeft":
		return r.unblockTVRepository.PressLeft(ctx)
	case "pressRight":
		return r.unblockTVRepository.PressRight(ctx)
	case "pressVolumeUp":
		return r.unblockTVRepository.PressVolumeUp(ctx)
	case "pressVolumeDown":
		return r.unblockTVRepository.PressVolumeDown(ctx)
	case "pressOK":
		return r.unblockTVRepository.PressOK(ctx)
	case "pressUp":
		return r.unblockTVRepository.PressUp(ctx)
	case "pressDown":
		return r.unblockTVRepository.PressDown(ctx)
	case "pressHome":
		return r.unblockTVRepository.PressHome(ctx)
	case "pressMenu":
		return r.unblockTVRepository.PressMenu(ctx)
	case "pressReturn":
		return r.unblockTVRepository.PressReturn(ctx)
	}
	return domain.ErrUnknownCommand
}
