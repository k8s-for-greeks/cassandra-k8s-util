# Copyright 2017 K8s For Greeks / Vorstella
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

UNAME := $(shell uname -s)
BUILD_NUMBER ?= "local"
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./*vendor/*")
NAME := cassandra-k8s-util

build: govet
	mkdir -p .bin/linux-amd64 .bin/darwin
	GOOS=darwin GOARCH=amd64 go build -o .bin/darwin/${NAME} ./

ifeq ($(UNAME),Linux)
	cp .bin/linux-amd64/${NAME} $(GOPATH)/bin/
else
	cp .bin/darwin/${NAME} $(GOPATH)/bin/
endif

build-linux:
	GOOS=linux GOARCH=amd64 go build -o .bin/linux-amd64/${NAME} ./

full-build: build build-linux

govet: 
	go vet \
	github.com/k8s-for-greeks/${NAME}/cmd/... \
	github.com/k8s-for-greeks/${NAME}/pkg/...

test:  
	go test -v github.com/k8s-for-greeks/${NAME}/pkg/... -args -v=1 -logtostderr
	go test -v github.com/k8s-for-greeks/${NAME}/cmd/... -args -v=1 -logtostderr

fmt:
	gofmt -w -s cmd/
	gofmt -w -s pkg/
	
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

#verify-boilerplate verify-gofmt verify-packages

ci: govet verify-gofmt full-build test
	echo "Done!"
