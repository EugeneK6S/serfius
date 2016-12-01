# Makefile
#
#


.PHONY: build clean install xc-build 

NAME := serfius
VERSION := ${shell cat .version}
BUILD_HASH := `git rev-parse --short HEAD`

LDFLAGS := "-X main.version=${VERSION} -X main.build=${BUILD_HASH}"


default: build

build: 
	@echo Building ${NAME}:${VERSION}
	go build -v -o ./bin/${NAME} -ldflags ${LDFLAGS} ./src/main.go
	@echo Finished

clean: 
	@echo Removing binaries and cleaning up my closet...
	@rm -rf bin/*

install: build
	install -m 755 ./bin/${NAME} /usr/bin/${NAME}