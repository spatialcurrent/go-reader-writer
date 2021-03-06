#!/bin/bash

# =================================================================
#
# Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export testdata_local=$(realpath "${DIR}/../testdata")

export "testdata_s3"="${GRW_TESTDATA_S3:-}"

_testDevice() {
  local algorithm=$1
  local expected='hello world'
  local output=$(echo 'hello world' | grw --output-compression $algorithm | grw --input-compression $algorithm)
  assertEquals "unexpected output" "${expected}" "${output}"
}

_testRead() {
  local algorithm=$1
  local filePath=$2
  local expected='hello world'
  local output=$(grw --input-compression ${algorithm} ${filePath} -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

#
# Test Reading/Writing Devices
#


testDevicePath() {
  local expected='hello world'
  local output=$(echo 'hello world' | grw /dev/stdin /dev/stdout)
  assertEquals "unexpected output" "${expected}" "${output}"
}

testDeviceName() {
  local expected='hello world'
  local output=$(echo 'hello world' | grw stdin stdout)
  assertEquals "unexpected output" "${expected}" "${output}"
}

testDeviceNone() {
  _testDevice 'none'
}

testDeviceGzip() {
  _testDevice 'gzip' - -
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

testDeviceDashes() {
  local expected='hello world'
  local output=$(echo 'hello world' | grw - -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

testDeviceAbsolutePath() {
  local expected='hello world'
  local output=$(echo 'hello world' | grw '/dev/stdin' '/dev/stdout')
  assertEquals "unexpected output" "${expected}" "${output}"
}

#
# Test Reading Local Files
#

testReadFileNone() {
  _testRead 'none' "${testdata_local}/doc.txt"
}

testReadFileBzip2() {
  _testRead 'bzip2' "${testdata_local}/doc.txt.bz2"
}

testReadFileGzip() {
  _testRead 'gzip' "${testdata_local}/doc.txt.gz"
}

testReadFileFlate() {
  _testRead 'flate' "${testdata_local}/doc.txt.f"
}

testReadFileSnappy() {
  _testRead 'snappy' "${testdata_local}/doc.txt.sz"
}

testReadFileZlib() {
  _testRead 'zlib' "${testdata_local}/doc.txt.z"
}

testReadFileZip() {
  _testRead 'zip' "${testdata_local}/doc.txt.zip"
}

#
# Test Splitting Input
#

testSplit() {
  local input='hello\nbeautiful\nworld'
  local expected='hello world'
  echo -e "${input}" | grw --split-lines 1 - "${SHUNIT_TMPDIR}/test_split_#.txt"
  local output=$(cat "${SHUNIT_TMPDIR}/test_split_1.txt" "${SHUNIT_TMPDIR}/test_split_2.txt" "${SHUNIT_TMPDIR}/test_split_3.txt")
  assertEquals "unexpected output" "$(echo -e "${output}")" "${output}"
}


#
# Test Reading S3 Objects
#

testReadS3None() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'none' "${testdata_s3}/doc.txt"
  else
    echo "* skipping"
  fi
}

testReadS3Bzip2() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'bzip2' "${testdata_s3}/doc.txt.bz2"
  else
    echo "* skipping"
  fi
}

testReadS3Gzip() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'gzip' "${testdata_s3}/doc.txt.gz"
  else
    echo "* skipping"
  fi
}

testReadS3Flate() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'flate' "${testdata_s3}/doc.txt.f"
  else
    echo "* skipping"
  fi
}

testReadS3Snappy() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'snappy' "${testdata_s3}/doc.txt.sz"
  else
    echo "* skipping"
  fi
}

testReadS3Zlib() {
  if [[ ! -z "${testdata_s3}" ]]; then
    _testRead 'zlib' "${testdata_s3}/doc.txt.z"
  else
    echo "* skipping"
  fi
}

testReadS3Zip() {
  if [[ ! -z "${testdata_s3}" ]]; then
   _testRead 'zip' "${testdata_s3}/doc.txt.zip"
  else
    echo "* skipping"
  fi
}

oneTimeSetUp() {
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
  echo "Reading testdata from ${testdata_local}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"