all: build test

build:
	go build -o bin/prism ./main.go

test:
	go test ./ ...

gen:
	go generate ./ ...
