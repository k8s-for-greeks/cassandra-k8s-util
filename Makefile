UNAME := $(shell uname -s)
BUILD_NUMBER ?= "local"
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./*vendor/*")
NAME := cassandra-k8s-util

build: 
	go vet -v
	mkdir -p .bin/linux-amd64 .bin/darwin
	#GOOS=linux GOARCH=amd64 go build -o .bin/linux-amd64/${NAME} ./
	GOOS=darwin GOARCH=amd64 go build -o .bin/darwin/${NAME} ./

ifeq ($(UNAME),Linux)
	cp .bin/linux-amd64/${NAME} $(GOPATH)/bin/
else
	cp .bin/darwin/${NAME} $(GOPATH)/bin/
endif


test: build
	go test -v github.com/k8s-for-greeks/${NAME}/pkg/... -args -v=1 -logtostderr
	go test -v github.com/k8s-for-greeks/${NAME}/cmd/... -args -v=1 -logtostderr

fmt:
	gofmt -w -s cmd/
	gofmt -w -s pkg/

copydeps:
	rsync -avz _vendor/ vendor/ --delete --exclude vendor/  --exclude .git
	
# --------------------------------------------------
# Continuous integration targets

verify-boilerplate:
	sh -c hack/verify-boilerplate.sh

.PHONY: verify-gofmt
verify-gofmt:
	hack/verify-gofmt.sh

.PHONY: verify-packages
verify-packages:
	hack/verify-packages.sh

ci: kops nodeup-gocode examples test govet verify-boilerplate verify-gofmt verify-packages
	echo "Done!"

