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

const base_https = process.env.GRW_TESTDATA_HTTPS;

const base_s3 = process.env.GRW_TESTDATA_S3;

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
    
    if(base_https != undefined) {
      describe('https', () => {

        it('read from a remote file over HTTPS', done => {
          read(base_https+"/doc.txt", "none", {"bufferSize": 4096}, callback(done));
        });
  
        it('read from a remote file over HTTPS and decompress using bzip2', done => {
          read(base_https+"/doc.txt.bz2", "bzip2", {"bufferSize": 4096}, callback(done));
        });
        
        it('read from a remote file over HTTPS and decompress using flate', done => {
          read(base_https+"/doc.txt.f", "flate", {"bufferSize": 4096}, callback(done));
        });
  
        it('read from a remote file over HTTPS and decompress using gzip', done => {
          read(base_https+"/doc.txt.gz", "gzip", {"bufferSize": 4096}, callback(done));
        });
  
        it('read from a remote file over HTTPS and decompress using snappy', done => {
          read(base_https+"/doc.txt.sz", "snappy", {"bufferSize": 4096}, callback(done));
        });
  
        it('read from a remote file over HTTPS and decompress using zip', done => {
          read(base_https+"/doc.txt.zip", "zip", {"bufferSize": 4096}, callback(done));
        });
        
        it('read from a remote file over HTTPS and decompress using zlib', done => {
          read(base_https+"/doc.txt.z", "zlib", {"bufferSize": 4096}, callback(done));
        });
  
      });
    }
    
    if(base_s3 != undefined) {
      describe('s3', () => {
  
        it('read from a file on S3', done => {
          read(base_s3+"/doc.txt", "none", options, callback(done));
        });
  
        it('read from a file on S3 and decompress using bzip2', done => {
          read(base_s3+"/doc.txt.bz2", "bzip2", options, callback(done));
        });
        
        it('read from a file on S3 and decompress using flate', done => {
          read(base_s3+"/doc.txt.f", "flate", options, callback(done));
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
        
        it('read from a file on S3 and decompress using zlib', done => {
          read(base_s3+"/doc.txt.z", "zlib", options, callback(done));
        });
  
      });
    }

  });

});
