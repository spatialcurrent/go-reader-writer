#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST
echo "******************"
echo "Building AAR for go-reader-writer (grw)"
cd $DEST
gomobile bind -target android -javapkg=com.spatialcurrent -o grw.aar github.com/spatialcurrent/go-reader-writer/grw
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-reader-writer (grw)"
    exit 1
fi
echo "Executable built at $DEST"
