linux:
	CGO_ENABLED=0
	GO_OS=linux
	GOARCH=amd64
	go build -o build/feature-flag-demo

windows:
	CGO_ENABLED=0
	GO_OS=windows
	GOARCH=amd64
	go build -o build/feature-flag-demo.exe
