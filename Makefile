PACKAGES = $(shell go list ./... | grep -v /vendor/)

default: build

build: check
	go build

check: vet lint errcheck interfacer test

vet:
	go vet $(PACKAGES)

lint:
	golint -set_exit_status $(PACKAGES)

errcheck:
	errcheck -ignore 'io:Close' $(PACKAGES)

interfacer:
	interfacer $(PACKAGES)

test:
	go test -cover $(PACKAGES)

clean:
	go clean

deps:
	go get -u github.com/FiloSottile/gvt
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/mvdan/interfacer/cmd/interfacer
