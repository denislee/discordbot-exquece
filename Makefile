fmt:
	go fmt ./...

build:
	go fmt ./... && \
	go get && go build

run:
	go fmt ./... && \
	go get && go build && ./$${PWD##*/}
