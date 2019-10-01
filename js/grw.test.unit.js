// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { TextDecoder } = require('text-encoding');
const { read, algorithms, schemes } = global.grw;

function log(str) {
  console.log(str.replace(/\n/g, "\\n").replace(/\t/g, "\\t").replace(/"/g, "\\\""));
}

describe('grw', () => {

  it('checks the available schemes', () => {
    expect(Array.isArray(schemes)).toBe(true);
    expect(schemes.sort()).toEqual(["file", "http", "https", "s3"]);
  });

  it('checks the available compression algorithms', () => {
    expect(Array.isArray(algorithms)).toBe(true);
    expect(algorithms.sort()).toEqual(["bzip2", "flate", "gzip", "none", "snappy", "zip", "zlib"]);
  });

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

    describe('local file', () => {

      it('read from a local file', done => {
        read(__dirname+"/../testdata/doc.txt", "none", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using bzip2', done => {
        read(__dirname+"/../testdata/doc.txt.bz2", "bzip2", {"bufferSize": 4096}, callback(done));
      });
      
      it('read from a local file and decompress using flate', done => {
        read(__dirname+"/../testdata/doc.txt.f", "flate", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using gzip', done => {
        read(__dirname+"/../testdata/doc.txt.gz", "gzip", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using snappy', done => {
        read(__dirname+"/../testdata/doc.txt.sz", "snappy", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using zip', done => {
        read(__dirname+"/../testdata/doc.txt.zip", "zip", {"bufferSize": 4096}, callback(done));
      });
      
      it('read from a local file and decompress using zlib', done => {
        read(__dirname+"/../testdata/doc.txt.z", "zlib", {"bufferSize": 4096}, callback(done));
      });

    });

  });

});
