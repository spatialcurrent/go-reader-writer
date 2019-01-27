# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build_so:
	go build -o grw.so -buildmode=c-shared github.com/spatialcurrent/go-reader-writer/plugins/grw
build_c:
	mkdir -p bin && gcc -o bin/grw_test_c test.c -L. -l:grw.so
build: build_so build_c
run:
	LD_LIBRARY_PATH=. bin/grw_test_c
clean:
	rm -f grw.h grw.so bin/grw_test_c
