[![Build Status](https://travis-ci.org/spatialcurrent/go-reader-writer.svg)](https://travis-ci.org/spatialcurrent/go-reader-writer) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-reader-writer)](https://goreportcard.com/report/spatialcurrent/go-reader-writer)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-reader-writer?status.svg)](https://godoc.org/github.com/spatialcurrent/go-reader-writer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-reader-writer/blob/master/LICENSE)

# go-reader-writer

# Description

**go-reader-writer** (aka GRW) is a simple library for managing reading/writing of resources.  GRW can read from `bzip2`, `gzip`, `snappy`, and `zip` resources and write to `gzip` and `snappy` resources.

Using cross compilers, this library can also be called by other languages.  This library is cross compiled into a Shared Object file (`*.so`).  The Shared Object file can be called by `C`, `C++`, and `Python` on Linux machines.  See the examples folder for patterns that you can use.  This library is also compiled to pure `JavaScript` using [GopherJS](https://github.com/gopherjs/gopherjs).

# Usage

**CLI**

You can use the command line tool to convert between formats.

```
Usage: grw [-input_uri INPUT_URI] [-input_compression [bzip2|gzip|snappy|zip|none]] [-output_uri OUTPUT_URI] [-output_compression [bzip2|gzip|snappy|zip|none]] [-verbose] [-version]
Options:
  -aws_access_key_id string
        Defaults to value of environment variable AWS_ACCESS_KEY_ID
  -aws_default_region string
        Defaults to value of environment variable AWS_DEFAULT_REGION.
  -aws_secret_access_key string
        Defaults to value of environment variable AWS_SECRET_ACCESS_KEY.
  -aws_session_token string
        Defaults to value of environment variable AWS_SESSION_TOKEN.
  -help
        Print help
  -input_buffer_size int
        the input reader buffer size (default 4096)
  -input_compression string
        Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.
  -input_uri string
        "stdin" or uri to input file (default "stdin")
  -output_append
        append output to resource
  -output_buffer_size int
        the output writer buffer size (default 4096)
  -output_compression string
        Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, zip, or none.
  -output_uri string
        "stdout" or uri to output resource (default "stdout")
  -version
        Prints version to stdout
```

**Go**

You can import **go-reader-writer** as a library with:

```go
import (
  "github.com/spatialcurrent/go-reader-writer/grw"
)
...
```

See [grw](https://godoc.org/github.com/spatialcurrent/go-reader-writer/grw) in GoDoc for information on how to use Go API.

**JavaScript**

```html
<html>
  <head>
    <script src="https://...grw.js"></script>
  </head>
  <body>
    <script>
      grw.open(uri, "none", function(text){ ... })
      ...
    </script>
  </body>
</html>
```

# Examples:

TBD

# Building

**CLI**

The command line go-reader-wrtier program can be built with the `scripts/build_cli.sh` script.

**JavaScript**

You can compile go-reader-writer to pure JavaScript with the `scripts/build_javascript.sh` script.

**Shared Object**

The `scripts/build_so.sh` script is used to build a Shared Object (`*.go`), which can be called by `C`, `C++`, and `Python` on Linux machines.

**Changing Destination**

The default destination for build artifacts is `go-reader/bin`, but you can change the destination with a CLI argument.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `bash scripts/build_cli.sh /usr/local/go/bin`

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-reader-writer/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
