# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build:
	go build -o grw.so -buildmode=c-shared github.com/spatialcurrent/go-reader-writer/plugins/grw
run:
	LD_LIBRARY_PATH=. python test.py
clean:
	rm -f grw.h grw.so
