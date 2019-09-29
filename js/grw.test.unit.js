// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

const { TextDecoder } = require('text-encoding');
const { read, algorithms, schemes } = global.grw;

const base_https = process.env.GRW_TESTDATA_HTTPS;

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
        read("./testdata/doc.txt", "none", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using bzip2', done => {
        read("./testdata/doc.txt.bz2", "bzip2", {"bufferSize": 4096}, callback(done));
      });
      
      it('read from a local file and decompress using flate', done => {
        read("./testdata/doc.txt.f", "flate", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using gzip', done => {
        read("./testdata/doc.txt.gz", "gzip", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using snappy', done => {
        read("./testdata/doc.txt.sz", "snappy", {"bufferSize": 4096}, callback(done));
      });

      it('read from a local file and decompress using zip', done => {
        read("./testdata/doc.txt.zip", "zip", {"bufferSize": 4096}, callback(done));
      });
      
      it('read from a local file and decompress using zlib', done => {
        read("./testdata/doc.txt.z", "zlib", {"bufferSize": 4096}, callback(done));
      });

    });
    
    if(GRW_TESTDATA_HTTPS != undefined) {
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

  });

});
