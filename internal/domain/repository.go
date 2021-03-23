package domain

import (
	"context"
)

type SerialPortRepository interface {
	Execute(ctx context.Context, device string, command string) error
}

type SharpTVRepository interface {
	PressPower(ctx context.Context) error
	PressInput(ctx context.Context) error
	PressLeft(ctx context.Context) error
	PressRight(ctx context.Context) error
	PressVolumeUp(ctx context.Context) error
	PressVolumeDown(ctx context.Context) error
	PressEnter(ctx context.Context) error
}

type LogiSpeakerRepository interface {
	PressPower(ctx context.Context) error
}

type UnblockTVRepository interface {
	PressPower(ctx context.Context) error
	PressUp(ctx context.Context) error
	PressDown(ctx context.Context) error
	PressLeft(ctx context.Context) error
	PressRight(ctx context.Context) error
	PressVolumeUp(ctx context.Context) error
	PressVolumeDown(ctx context.Context) error
	PressOK(ctx context.Context) error
	PressHome(ctx context.Context) error
	PressMenu(ctx context.Context) error
	PressReturn(ctx context.Context) error
}
