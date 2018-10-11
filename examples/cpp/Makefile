# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build_so:
	go build -o grw.so -buildmode=c-shared github.com/spatialcurrent/go-reader-writer/plugins/grw
build_cpp:
	mkdir -p bin && g++ -o bin/grw_test_cpp test.cpp -L. -l:grw.so
build: build_so build_cpp
run:
	LD_LIBRARY_PATH=. bin/grw_test_cpp
clean:
	rm -f grw.h grw.so bin/grw_test_cpp
