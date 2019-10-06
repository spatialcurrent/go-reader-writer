[![CircleCI](https://circleci.com/gh/spatialcurrent/go-reader-writer/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-reader-writer/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-reader-writer)](https://goreportcard.com/report/spatialcurrent/go-reader-writer)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-reader-writer?status.svg)](https://godoc.org/github.com/spatialcurrent/go-reader-writer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-reader-writer/blob/master/LICENSE)

# go-reader-writer

## Description

**go-reader-writer** (aka GRW) is a simple library for reading and writing compressed resources by uri.  GRW supports the following compression algorithms.

| Algorithm | Read |  Write |
| ---- | ------ |  ------ |
| bzip2 | ✓ | - |
| flate | ✓ | ✓ |
| gzip | ✓ | ✓ |
| snappy | ✓ | ✓ |
| zip | ✓ | - |
| zlib | ✓ | ✓ |

Using cross compilers, this library can also be called by other languages.  This library is cross compiled into a Shared Object file (`*.so`).  The Shared Object file can be called by `C`, `C++`, and `Python` on Linux machines.  See the examples folder for patterns that you can use.  This library is also compiled to pure `JavaScript` using [GopherJS](https://github.com/gopherjs/gopherjs).

## Usage

**CLI**

The command line tool, `grw`, can be used to easily read and write compressed resources by uri.  We currently support the following platforms.

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

Pull requests to support other platforms are welcome!  See the [CLI.md](docs/CLI.md) document for detailed usage and examples.

**Go**

You can install the go-reader-writer packages with.


```shell
go get -u -d github.com/spatialcurrent/go-reader-writer/...
```

See [grw](https://godoc.org/github.com/spatialcurrent/go-reader-writer/pkg/grw) in GoDoc for information on how to use Go API.

**Node**

GRW is built as a module.  In modern JavaScript, the module can be imported using [destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment).

```javascript
const { read, algorithms, schemes } = require('./dist/grw.mod.min.js');
```

In legacy JavaScript, you can use the `grw.global.js` file that simply adds `grw` to the global scope.

**C**

A variant of the reader and writer functions are exported in a Shared Object file (`*.so`), which can be called by `C`, `C++`, and `Python` programs on Linux machines.  For complete patterns for `C`, `C++`, and `Python`, see the `examples` folder in this repo.

## Releases

**go-reader-writer** is currently in **alpha**.  Find releases at [https://github.com/spatialcurrent/go-reader-writer/releases](https://github.com/spatialcurrent/go-reader-writer/releases).  See the **Building** section below to build from scratch.

**JavaScript**

- `grw.global.js`, `grw.global.js.map` - JavaScript global build  with source map
- `grw.global.min.js`, `grw.global.min.js.map` - Minified JavaScript global build with source map
- `grw.mod.js`, `grw.mod.js.map` - JavaScript module build  with source map
- `grw.mod.min.js`, `grw.mod.min.js.map` - Minified JavaScript module with source map

**Darwin**

- `grw_darwin_amd64` - CLI for Darwin on amd64 (includes `macOS` and `iOS` platforms)

**Linux**

- `grw_linux_amd64` - CLI for Linux on amd64
- `grw_linux_amd64` - CLI for Linux on arm64
- `grw_linux_amd64.h`, `grw_linuxamd64.so` - Shared Object for Linux on amd64
- `grw_linux_armv7.h`, `grw_linux_armv7.so` - Shared Object for Linux on ARMv7
- `grw_linux_armv8.h`, `grw_linux_armv8.so` - Shared Object for Linux on ARMv8

**Windows**

- `grw_windows_amd64.exe` - CLI for Windows on amd64

## Examples

For the CLI, see the examples in the [CLI.md](docs/CLI.md) document.

For Go, see the examples in the [grw](https://godoc.org/github.com/spatialcurrent/go-reader-writer/pkg/grw) package documentation.

## Building

Use `make help` to see help information for each target.

**CLI**

The `make build_cli` script is used to build executables for Linux and Windows.

**JavaScript**

You can compile GRW to pure JavaScript with the `make build_javascript` script.

**Shared Object**

The `make build_so` script is used to build Shared Objects (`*.so`), which can be called by `C`, `C++`, and `Python` on Linux machines.

**Changing Destination**

The default destination for build artifacts is `go-reader-writer/bin`, but you can change the destination with an environment variable.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `DEST=/usr/local/go/bin make build_cli`

## Testing

**CLI**

To run CLI testes use `make test_cli`, which uses [shUnit2](https://github.com/kward/shunit2).  If you recive a `shunit2:FATAL Please declare TMPDIR with path on partition with exec permission.` error, you can modify the `TMPDIR` environment variable in line or with `export TMPDIR=<YOUR TEMP DIRECTORY HERE>`. For example:

```
TMPDIR="/usr/local/tmp" make test_cli
```

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

**JavaScript**

To run JavaScript tests, first install [Jest](https://jestjs.io/) using `make deps_javascript`, use [Yarn](https://yarnpkg.com/en/), or another method.  Then, build the JavaScript module with `make build_javascript`.  To run tests, use `make test_javascript`.  You can also use the scripts in the `package.json`.

## Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-reader-writer/blob/master/CONTRIBUTING.md) for how to get started.

## License

This work is distributed under the **MIT License**.  See **LICENSE** file.
