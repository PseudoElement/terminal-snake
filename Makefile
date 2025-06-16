build-mac:
	go build -o ./builds/snake-mac

build-win64:
	GOOS=windows GOARCH=amd64 go build -o builds/snake-win64.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o builds/snake-linux

run:
	make build-mac && ./builds/snake-mac