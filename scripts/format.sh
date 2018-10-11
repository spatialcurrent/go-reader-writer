#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "******************"
echo "Formatting $DIR/../grw"
cd $DIR/../grw
go fmt
echo "Formatting $DIR/../plugins/grw"
cd $DIR/../plugins/grw
go fmt
echo "Formatting $DIR/../cmd/grw"
cd $DIR/../cmd/grw/
go fmt
echo "Formatting $DIR/../cmd/grw.js"
cd $DIR/../cmd/grw.js
go fmt