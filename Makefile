# =================================================================
#
# Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

ifdef GOPATH
GCFLAGS=-trimpath=$(shell printenv GOPATH)/src
else
GCFLAGS=-trimpath=$(shell go env GOPATH)/src
endif

LDFLAGS=-X main.gitBranch=$(shell git branch | grep \* | cut -d ' ' -f2) -X main.gitCommit=$(shell git rev-list -1 HEAD)

.PHONY: help
help:  ## Print the help documentation
	@grep -E '^[a-zA-Z0-9_-\]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#
# Dependencies
#

deps_arm:  ## Install dependencies to cross-compile to ARM
	apt-get install -y gcc-arm-linux-gnueabi g++-arm-linux-gnueabi
	# ARMv7
	#apt-get install -y libc6-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev gcc-arm-linux-gnueabi g++-arm-linux-gnueabi

deps_arm64:  ## Install dependencies to cross-compile to ARM
	apt-get install -y gcc-aarch64-linux-gnu g++-aarch64-linux-gnu

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

bin/gobind:
	go build -o bin/gobind golang.org/x/mobile/cmd/gobind

bin/goimports:
	go build -o bin/goimports golang.org/x/tools/cmd/goimports

bin/gomobile:
	go build -o bin/gomobile golang.org/x/mobile/cmd/gomobile

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
	CGO_ENABLED=1 go build -o bin/grw.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw

bin/grw_linux_386.so:  ## Compile Shared Object for Linux / 386
	scripts/build-so linux 386

bin/grw_linux_amd64.so:  ## Compile Shared Object for Linux / amd64
	# https://golang.org/cmd/link/
	# CGO Enabled : https://github.com/golang/go/issues/24068
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o bin/grw_linux_amd64.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw
	scripts/build-so linux amd64

bin/grw_linux_arm_v7.so:  ## Compile Shared Object for Linux / ARMv7
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	#GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -ldflags "-linkmode external -extldflags -static" -o bin/grw_linux_armv7.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw
	scripts/build-so linux arm 7

bin/grw_linux_arm64.so:   ## Compile Shared Object for Linux / ARMv8
	# LDFLAGS - https://golang.org/cmd/link/
	# CGO Enabled  - https://github.com/golang/go/issues/24068
	# GOARM/GOARCH Compatability Table - https://github.com/golang/go/wiki/GoArm
	# ARM Cross Compiler Required - https://www.acmesystems.it/arm9_toolchain
	# Dependencies - https://www.96boards.org/blog/cross-compile-files-x86-linux-to-96boards/
	#GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -ldflags "-linkmode external -extldflags -static" -o bin/grw_linux_armv8.so -buildmode=c-shared -ldflags "$(LDFLAGS)" -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/plugins/grw
	scripts/build-so linux arm64


.PHONY: build_so
build_so: bin/grw_linux_amd64.so bin/grw_linux_arm_v7.so bin/grw_linux_arm64.so  ## Build Shared Objects (.so)

#
# Android
#

bin/grw.aar: bin/gobind bin/gomobile   ## Build Android Archive Library
	bin/gomobile init
	bin/gomobile bind -target android -javapkg=com.spatialcurrent -o bin/grw.aar -gcflags="$(GCFLAGS)" github.com/spatialcurrent/go-reader-writer/pkg/android

build_android: bin/grw.aar  ## Build artifacts for Android

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

run_example_python3: bin/grw.so  ## Run Python example
	LD_LIBRARY_PATH=bin python3 examples/python/test.py

## Clean

.PHONY: clean
clean:  ## Clean artifacts
	rm -fr bin
