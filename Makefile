mkfile_path := $(abspath .)

default: build

test:
	go test ./...

build: test
	go build -o build/go-db2-example.exe

build_linux: test
	go build -o build/go-db2-example

build-container:
	docker build -t go-db2-build .
	docker run --rm --name go-db2-build -v $(mkfile_path)/build:/go/src/app/build go-db2-build
