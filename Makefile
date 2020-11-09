# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

ifdef GOPATH
GCFLAGS=-trimpath=$(shell printenv GOPATH)/src
else
GCFLAGS=-trimpath=$(shell go env GOPATH)/src
endif

LDFLAGS=-X main.gitBranch=$(shell git branch | grep \* | cut -d ' ' -f2) -X main.gitCommit=$(shell git rev-list -1 HEAD)

ifndef DEST
DEST=bin
endif

.PHONY: help
help:  ## Print the help documentation
	@grep -E '^[a-zA-Z0-9_-\]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#
# Dependencies
#

deps_go:  ## Install Go dependencies
	go get -d -t ./...
	go get -d honnef.co/go/js/xhr # used in JavaScript build

.PHONY: deps_go_test
deps_go_test: ## Download Go dependencies for tests
	go get golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # download shadow
	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow # install shadow
	go get -u github.com/kisielk/errcheck # download and install errcheck
	go get -u github.com/client9/misspell/cmd/misspell # download and install misspell
	go get -u github.com/gordonklaus/ineffassign # download and install ineffassign
	go get -u honnef.co/go/tools/cmd/staticcheck # download and instal staticcheck
	go get -u golang.org/x/tools/cmd/goimports # download and install goimports

deps_arm:  ## Install dependencies to cross-compile to ARM
	# ARMv7
	apt-get install -y libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev gcc-arm-linux-gnueabi g++-arm-linux-gnueabi
  # ARMv8
	apt-get install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu

deps_gopherjs:  ## Install GopherJS
	go get -u github.com/gopherjs/gopherjs

deps_javascript:  ## Install dependencies for JavaScript tests
	npm install .

#
# Go building, formatting, testing, and installing
#

fmt:  ## Format Go source code
	go fmt $$(go list ./... )

.PHONY: imports
imports: bin/goimports ## Update imports in Go source code
	# If missing, install goimports with: go get golang.org/x/tools/cmd/goimports
	bin/goimports -w -local github.com/spatialcurrent/go-reader-writer,github.com/spatialcurrent $$(find . -iname '*.go')

vet: ## Vet Go source code
	go vet github.com/spatialcurrent/go-reader-writer/pkg/... # vet packages
	go vet github.com/spatialcurrent/go-reader-writer/cmd/... # vet commands
	go vet github.com/spatialcurrent/go-reader-writer/plugins/... # vet plugins

tidy: ## Tidy Go source code
	go mod tidy

.PHONY: test_go
test_go: bin/errcheck bin/ineffassign bin/misspell bin/staticcheck bin/shadow ## Run Go tests
	bash scripts/test.sh

.PHONY: test_cli
test_cli: bin/grw ## Run CLI tests
	bash scripts/test-cli.sh

install:  ## Install grw CLI on current platform
	go install github.com/spatialcurrent/go-reader-writer/cmd/grw

#.PHONY: build
#build: build_cli build_javascript build_so build_android  ## Build CLI, Shared Objects (.so), JavaScript, and Android

#
# Command line Programs
#

bin/errcheck:
	go build -o bin/errcheck github.com/kisielk/errcheck

bin/goimports:
	go build -o bin/goimports golang.org/x/tools/cmd/goimports

bin/gox:
	go build -o bin/gox github.com/mitchellh/gox

bin/ineffassign:
	go build -o bin/ineffassign github.com/gordonklaus/ineffassign

bin/misspell:
	go build -o bin/misspell github.com/client9/misspell/cmd/misspell

bin/staticcheck:
	go build -o bin/staticcheck honnef.co/go/tools/cmd/staticcheck

bin/shadow:
	go build -o bin/shadow golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

bin/grw: ## Build grw CLI for Darwin / amd64
	go build -o bin/grw github.com/spatialcurrent/go-reader-writer/cmd/grw

bin/grw_linux_amd64: bin/gox ## Build grw CLI for Darwin / amd64
	scripts/build-release linux amd64

.PHONY: build
build: bin/grw

.PHONY: build_release
build_release: bin/gox
	scripts/build-release

#
# Shared Objects
#

bin/grw.so:  ## Compile Shared Object for current platform
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	CGO_ENABLED=1 go build -o $(DEST)/grw.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw

bin/grw_linux_amd64.so:  ## Compile Shared Object for Linux / amd64
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(DEST)/grw_linux_amd64.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw

bin/grw_linux_armv7.so:  ## Compile Shared Object for Linux / ARMv7
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/grw_linux_armv7.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw

bin/grw_linux_armv8.so:   ## Compile Shared Object for Linux / ARMv8
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	# Dependencies - https://www.96boards.org/blog/cross-compile-files-x86-linux-to-96boards/
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -ldflags "-linkmode external -extldflags -static" -o $(DEST)/grw_linux_armv8.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw

.PHONY: build_so
build_so: bin/grw_linux_amd64.so bin/grw_linux_armv7.so bin/grw_linux_armv8.so  ## Build Shared Objects (.so)

#
# Android
#

bin/grw.aar:  ## Build Android Archive Library
	gomobile bind -target android -javapkg=com.spatialcurrent -o $(DEST)/grw.aar -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/pkg/grw

build_android: bin/grw.arr  ## Build artifacts for Android

#
# JavaScript
#

dist/grw.mod.js:  ## Build JavaScript module
	gopherjs build -o dist/grw.mod.js github.com/spatialcurrent/go-reader-writer/cmd/grw.mod.js

dist/grw.mod.min.js:  ## Build minified JavaScript module
	gopherjs build -m -o dist/grw.mod.min.js github.com/spatialcurrent/go-reader-writer/cmd/grw.mod.js

dist/grw.global.js:  ## Build JavaScript library that attaches to global or window.
	gopherjs build -o dist/grw.global.js github.com/spatialcurrent/go-reader-writer/cmd/grw.global.js

dist/grw.global.min.js:  ## Build minified JavaScript library that attaches to global or window.
	gopherjs build -m -o dist/grw.global.min.js github.com/spatialcurrent/go-reader-writer/cmd/grw.global.js

build_javascript: dist/grw.mod.js dist/grw.mod.min.js dist/grw.global.js dist/grw.global.min.js  ## Build artifacts for JavaScript

test_javascript:  ## Run JavaScript tests
	npm run test

lint:  ## Lint JavaScript source code
	npm run lint

#
# Examples
#

bin/grw_example_c: bin/grw.so  ## Build C example
	mkdir -p bin && cd bin && gcc -o grw_example_c -I. ./../examples/c/main.c -L. -l:grw.so

bin/grw_example_cpp: bin/grw.so  ## Build C++ example
	mkdir -p bin && cd bin && g++ -o grw_example_cpp -I . ./../examples/cpp/main.cpp -L. -l:grw.so

run_example_c: bin/grw.so bin/grw_example_c  ## Run C example
	cd bin && LD_LIBRARY_PATH=. ./grw_example_c

run_example_cpp: bin/grw.so bin/grw_example_cpp  ## Run C++ example
	cd bin && LD_LIBRARY_PATH=. ./grw_example_cpp

run_example_python: bin/grw.so  ## Run Python example
	LD_LIBRARY_PATH=bin python examples/python/test.py

run_example_javascript: dist/grw.mod.js  ## Run JavaScript module example
	npm run examples

## Clean

.PHONY: clean
clean:  ## Clean artifacts
	rm -fr bin
	rm -fr dist
