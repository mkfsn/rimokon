package domain

import (
	"errors"
)

var (
	ErrUnknownDevice  = errors.New("unknown device")
	ErrUnknownCommand = errors.New("unknown command")
)
