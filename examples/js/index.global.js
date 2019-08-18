// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Used for file path logic
const path = require('path');

// TextDecoder is a polyfill for NodeJS to decode utf-8 bytes.
const { TextDecoder } = require('text-encoding');

// Loads grw into global scope
require('./../../dist/grw.global.min.js');

// Destructure grw into consts
const { read, schemes, algorithms } = global.grw;

console.log('Schemes:', schemes);
console.log();
console.log('Algorithms:', algorithms);
console.log();
console.log("************************************");
console.log();

read(path.join(__dirname, "../../testdata/doc.txt"), "none", {}, function(reader, err) {
  var { data, err } = reader.ReadAllAndClose();
  console.log("Error:", err);
  if (err != null) {
    return;
  }
  let str = new TextDecoder("utf-8").decode(Uint8Array.from(data));
  console.log('Output:');
  console.log(str);
  console.log();
  console.log("************************************");
  console.log();
});
