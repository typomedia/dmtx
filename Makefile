icon:
	rsrc -ico go.ico

build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go run .

clean:
	go clean

compile:
	go mod tidy
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/dmtx-linux-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/dmtx-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/dmtx-windows-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/dmtx-macos-amd64 .

all: clean compile
