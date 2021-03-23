package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mkfsn/rimokon/internal/domain"
	"go.bug.st/serial"
)

type serialPortRepository struct {
	port serial.Port
}

func NewSerialPortRepository(serialPort string, baudRate int) (domain.SerialPortRepository, error) {
	serialMode := &serial.Mode{
		BaudRate: baudRate,
		DataBits: 8,
	}

	port, err := serial.Open(serialPort, serialMode)
	if err != nil {
		return nil, err
	}

	return &serialPortRepository{port: port}, nil
}

type Request struct {
	Device  string `json:"device"`
	Command string `json:"command"`
}

type Response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func (s *serialPortRepository) Execute(ctx context.Context, device string, command string) error {
	request := Request{Device: device, Command: command}
	log.Printf("request: %#v\n", request)

	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = s.port.Write(b)
	if err != nil {
		return err
	}

	data, err := s.readUntil(ctx, '\n')
	if err != nil {
		return err
	}
	log.Printf("data: %s\n", data)

	var response Response

	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	log.Printf("response: %#v\n", response)

	if response.OK {
		return nil
	}

	return fmt.Errorf("%s", response.Error)
}

func (s *serialPortRepository) readUntil(ctx context.Context, b byte) ([]byte, error) {
	var data []byte

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		buf := make([]byte, 128)

		n, err := s.port.Read(buf)
		if err != nil {
			return nil, err
		}

		if n == 0 {
			break
		}

		data = append(data, buf[:n]...)

		if buf[n-1] == b {
			break
		}
	}

	return data, nil
}
