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
# Test Reading S3 Objects
#

if [[ ! -z "${testdata_s3}" ]]; then
  
  testReadS3None() {
    _testRead 'none' "${testdata_s3}/doc.txt"
  }
  
  testReadS3Bzip2() {
    _testRead 'bzip2' "${testdata_s3}/doc.txt.bz2"
  }
  
  testReadS3Gzip() {
    _testRead 'gzip' "${testdata_s3}/doc.txt.gz"
  }
  
  testReadS3Flate() {
    _testRead 'flate' "${testdata_s3}/doc.txt.f"
  }
  
  testReadS3Snappy() {
    _testRead 'snappy' "${testdata_s3}/doc.txt.sz"
  }
  
  testReadS3Zlib() {
    _testRead 'zlib' "${testdata_s3}/doc.txt.z"
  }
  
  testReadS3Zip() {
    _testRead 'zip' "${testdata_s3}/doc.txt.zip"
  }
  
fi

oneTimeSetUp() {
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
  echo "Reading testdata from ${testdata_local}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"