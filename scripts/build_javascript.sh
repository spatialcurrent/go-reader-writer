#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Building Javascript artifacts for go-reader-writer"
cd $DEST
gopherjs build -o grw.js github.com/spatialcurrent/go-reader-writer/cmd/grw.js
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript artificats for go-reader-writer"
    exit 1
fi
gopherjs build -m -o grw.min.js github.com/spatialcurrent/go-reader-writer/cmd/grw.js
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript artificats for go-reader-writer"
    exit 1
fi
echo "JavaScript artificats built at $DEST"
