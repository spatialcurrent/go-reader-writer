// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const fs = require('fs');
const path = require("path");
const { TextDecoder } = require('text-encoding');

const { read, algorithms, schemes } = global.grw;

const base_s3 = process.env.GRW_TESTDATA_S3
expect(base_s3).toBeDefined();

var options = {
  "AWS_DEFAULT_REGION": process.env.AWS_DEFAULT_REGION,
  "AWS_REGION": process.env.AWS_REGION,
  "AWS_ACCESS_KEY_ID": process.env.AWS_ACCESS_KEY_ID,
  "AWS_SECRET_ACCESS_KEY": process.env.AWS_SECRET_ACCESS_KEY,
  "AWS_SESSION_TOKEN": process.env.AWS_SESSION_TOKEN
}

function log(str) {
  console.log(str.replace(/\n/g, "\\n").replace(/\t/g, "\\t").replace(/"/g, "\\\""));
}

describe('grw', () => {

  describe('reader', () => {

    var callback = function(done) {
      return function(reader, err){
        expect(err).toBeNull();
        expect(reader).toBeDefined();
        var { data, err } = reader.ReadAllAndClose();
        expect(err).toBeNull();
        let str = new TextDecoder("utf-8").decode(Uint8Array.from(data));
        expect(str).toEqual("hello world");
        done();
      };
    };

    describe('s3', () => {

      it('read from a file on S3', done => {
        read(base_s3+"/doc.txt", "none", options, callback(done));
      });

      it('read from a file on S3 and decompress using bzip2', done => {
        read(base_s3+"/doc.txt.bz2", "bzip2", options, callback(done));
      });

      it('read from a file on S3 and decompress using gzip', done => {
        read(base_s3+"/doc.txt.gz", "gzip", options, callback(done));
      });

      it('read from a file on S3 and decompress using snappy', done => {
        read(base_s3+"/doc.txt.sz", "snappy", options, callback(done));
      });

      it('read from a file on S3 and decompress using zip', done => {
        read(base_s3+"/doc.txt.zip", "zip", options, callback(done));
      });

    });

  });

});
