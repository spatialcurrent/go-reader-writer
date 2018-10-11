#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Building program for go-reader-writer (GRW)"
cd $DEST
for GOOS in darwin linux windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "grw_${GOOS}_amd64" github.com/spatialcurrent/go-reader-writer/cmd/grw
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-reader-writer (GRW)"
    exit 1
fi
echo "Executables built at $DEST"
