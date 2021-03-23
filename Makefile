.PHONY: all build clean upload

APP=rimokon

all:
	@echo "What do you want to do?"

build:
	GOOS=linux GOARCH=mipsle GOMIPS=hardfloat CGO_ENABLED=0 go build -v ./cmd/${APP}

clean:
	go clean

upload:
	scp ./${APP} root@mylinkit.local:./
