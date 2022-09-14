bl_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/simple-web-linux
