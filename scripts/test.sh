#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
set -eu
echo "******************"
echo "Running unit tests"
cd $DIR/../grw
go test
echo "******************"
echo "Using gometalinter with misspell, vet, ineffassign, and gosec"
echo "Testing $DIR/../grw"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../grw

echo "Testing $DIR/../plugins/grw"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../plugins/grw

echo "Testing $DIR/../cmd/grw"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../cmd/grw/

echo "Testing $DIR/../cmd/grw.js"
gometalinter --misspell-locale=US --disable-all --enable=misspell --enable=vet --enable=ineffassign --enable=gosec $DIR/../cmd/grw.js