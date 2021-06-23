BUILDDIR = bin
SERVICE_NAME = service
SERVER_NAME = server

VERSION = `git describe --tags --always --dirty`
BUILD = `date +%Y-%m-%d\ %H:%M`

VERSION := $(shell git describe --tags --always --dirty)
BUILD := $(shell date +%Y-%m-%d\ %H:%M)

# LDFLAGS=-ldflags="-X 'main.Version=${VERSION}' -X 'main.Build=${BUILD}' -X 'libcommon.Build=${BUILD}'"
LDFLAGS=-ldflags="-w -s -X 'gotest/libcommon.Version=${VERSION}' -X 'gotest/libcommon.Build=${BUILD}'"

hello:
	echo "Hello ${LDFLAGS}"
    echo ${LDFLAGS}

install:
	go install ./...
	go build ${LDFLAGS} -o $(BUILDDIR)/ ./...

clean:
	if [ -d $(BUILDDIR) ] ; then rm -rf $(BUILDDIR) ; fi

compile:
	# 64-Bit MacOS
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o $(BUILDDIR)/$(SERVICE_NAME)-darwin-amd64 .
	# 64-Bit Linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o $(BUILDDIR)/$(SERVICE_NAME)-linux-amd64 .

.PHONY: clean install
