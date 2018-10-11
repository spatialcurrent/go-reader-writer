#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Building Shared Object (*.so) for go-reader-writer"
cd $DEST
go build -o grw.so -buildmode=c-shared github.com/spatialcurrent/go-reader-writer/plugins/grw
if [[ "$?" != 0 ]] ; then
    echo "Error Building Shared Object (*.so) for go-reader-writer"
    exit 1
fi
echo "Shared Object (*.so) built at $DEST"
