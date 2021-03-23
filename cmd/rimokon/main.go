package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mkfsn/rimokon/internal/adapter"
	"github.com/mkfsn/rimokon/internal/repository"
	"github.com/mkfsn/rimokon/internal/usecase"
)

type Config struct {
	Host       string
	Port       int
	SerialPort string
	BaudRate   int
}

func main() {
	var config Config
	flag.StringVar(&config.Host, "host", "0.0.0.0", "http server host")
	flag.IntVar(&config.Port, "port", 8080, "http server port")
	flag.StringVar(&config.SerialPort, "serial-port", "/dev/ttyS0", "serial port")
	flag.IntVar(&config.BaudRate, "baud-rate", 57600, "baud rate")
	flag.Parse()

	serialPortRepository, err := repository.NewSerialPortRepository(config.SerialPort, config.BaudRate)
	if err != nil {
		log.Fatal(err)
	}

	rimokonUsecase := usecase.NewRimokonUsecase(
		repository.NewSharpTVRepository(serialPortRepository),
		repository.NewLogiSpeakerRepository(serialPortRepository),
		repository.NewUnblockTVRepository(serialPortRepository),
	)

	rimokonAdapter := adapter.NewHandler(rimokonUsecase)

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	server := http.Server{
		Addr:         addr,
		Handler:      rimokonAdapter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("listening on %s ...\n", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
