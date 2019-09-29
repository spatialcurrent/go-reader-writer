#!/bin/bash

# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export testdata=$(realpath "${DIR}/../testdata")

_testDevice() {
  local algorithm=$1
  local expected='hello world'
  local output=$(echo 'hello world' | grw --output-compression $algorithm | grw --input-compression $algorithm)
  assertEquals "unexpected output" "${expected}" "${output}"
}

_testReadFile() {
  local algorithm=$1
  local filePath=$2
  local expected='hello world'
  local output=$(grw --input-compression ${algorithm} ${filePath} -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

#
# Test Reading/Writing Devices
#

testDeviceNone() {
  _testDevice 'none'
}

testDeviceGzip() {
  _testDevice 'gzip'
}

testDeviceFlate() {
  _testDevice 'flate'
}

testDeviceSnappy() {
  _testDevice 'snappy'
}

testDeviceZlib() {
  _testDevice 'zlib'
}

#
# Test Reading Files
#

testReadFileNone() {
  _testReadFile 'none' "${testdata}/doc.txt"
}

testReadFileBzip2() {
  _testReadFile 'bzip2' "${testdata}/doc.txt.bz2"
}

testReadFileGzip() {
  _testReadFile 'gzip' "${testdata}/doc.txt.gz"
}

testReadFileFlate() {
  _testReadFile 'flate' "${testdata}/doc.txt.f"
}

testReadFileSnappy() {
  _testReadFile 'snappy' "${testdata}/doc.txt.sz"
}

testReadFileZlib() {
  _testReadFile 'zlib' "${testdata}/doc.txt.z"
}

testReadFileZip() {
  _testReadFile 'zip' "${testdata}/doc.txt.zip"
}

oneTimeSetUp() {
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
  echo "Reading testdata from ${testdata}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"