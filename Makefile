all: build test

build:
	go build -o bin/prism ./cmd/main.go

test:
	go test ./ ...

gen:
	go generate ./ ...
