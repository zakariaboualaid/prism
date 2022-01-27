all: build test

build:
	go build -o bin/prism ./main.go

run:
	./bin/prism -configFile=/Users/zak/Projects/prism/config-test.yaml

buildAndRun: build run

test:
	go test ./ ...

gen:
	go generate ./ ...
