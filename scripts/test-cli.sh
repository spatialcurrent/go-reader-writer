#!/bin/bash

# =================================================================
#
# Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

#export testdata_local=$(realpath "${DIR}/../testdata")
export testdata_local="${DIR}/../testdata"

export "testdata_s3"="${GRW_TESTDATA_S3:-}"

export "testdata_sftp"="${GRW_TESTDATA_SFTP:-}"

rand() {
  head -c $((${1}/2)) < /dev/random | base64
}

_testDevice() {
  local algorithm=$1
  local expected='hello world'
  local output=$(echo 'hello world' | "${DIR}/../bin/grw" --output-compression $algorithm | "${DIR}/../bin/grw" --input-compression $algorithm)
  assertEquals "unexpected output" "${expected}" "${output}"
}

_testRead() {
  local algorithm=$1
  local filePath=$2
  local expected='hello world'
  local output=$("${DIR}/../bin/grw" --input-compression ${algorithm} ${filePath} -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

_testWriteRead() {
  local algorithm=$1
  local filePath=$2
  local expected=$(rand 4096)
  echo "${expected}" | "${DIR}/../bin/grw" --output-compression "${algorithm}" - "${filePath}"
  local output=$("${DIR}/../bin/grw" --input-compression "${algorithm}" "${filePath}" -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

#
# Test Help
#

testHelp() {
  "${DIR}/../bin/grw" --help > /dev/null
}

#
# Test Reading/Writing Devices
#


testDevicePath() {
  local expected='hello world'
  local output=$(echo 'hello world' | "${DIR}/../bin/grw" file:///dev/stdin file:///dev/stdout)
  assertEquals "unexpected output" "${expected}" "${output}"
}

testDeviceName() {
  local expected='hello world'
  #local output=$(echo 'hello world' | "${DIR}/../bin/grw" stdin stdout)
  #assertEquals "unexpected output" "${expected}" "${output}"
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
  local output=$(echo 'hello world' | "${DIR}/../bin/grw" - -)
  assertEquals "unexpected output" "${expected}" "${output}"
}

testDeviceAbsolutePath() {
  local expected='hello world'
  local output=$(echo 'hello world' | "${DIR}/../bin/grw" 'file:///dev/stdin' 'file:///dev/stdout')
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
  echo -e "${input}" | "${DIR}/../bin/grw" --split-lines 1 - "${SHUNIT_TMPDIR}/test_split_#.txt"
  local output=$(cat "${SHUNIT_TMPDIR}/test_split_1.txt" "${SHUNIT_TMPDIR}/test_split_2.txt" "${SHUNIT_TMPDIR}/test_split_3.txt")
  assertEquals "unexpected output" "$(echo -e "${output}")" "${output}"
}

#
# Test Reading SFTP Objects
#

testWriteReadSFTPNone() {
  if [[ ! -z "${testdata_sftp}" ]]; then
    _testWriteRead 'none' "${testdata_sftp}/doc.txt"
  else
    echo "* skipping"
  fi
}

testWriteReadSFTPGZIP() {
  if [[ ! -z "${testdata_sftp}" ]]; then
    _testWriteRead 'gzip' "${testdata_sftp}/doc.txt.gz"
  else
    echo "* skipping"
  fi
}

testWriteReadSFTPFlate() {
  if [[ ! -z "${testdata_sftp}" ]]; then
    _testWriteRead 'flate' "${testdata_sftp}/doc.txt.f"
  else
    echo "* skipping"
  fi
}

testWriteReadSFTPSnappy() {
  if [[ ! -z "${testdata_sftp}" ]]; then
    _testWriteRead 'snappy' "${testdata_sftp}/doc.txt.sz"
  else
    echo "* skipping"
  fi
}

testWriteReadSFTPZlib() {
  if [[ ! -z "${testdata_sftp}" ]]; then
    _testWriteRead 'zlib' "${testdata_sftp}/doc.txt.z"
  else
    echo "* skipping"
  fi
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
