# CLI

- [Usage](#usage) - basic usage
- [Algorithms](#algorithms) - list of supported compression algorithms
- [Platforms](#platforms) - list of supported platforms
- [Releases](#releases) - where to find an executable
- [Examples](#examples)  - detailed usage exampels
- [Building](#building) - how to build the CLI
- [Testing](#testing) - test the CLI
- [Troubleshooting](#Troubleshooting) - how to troubleshoot common errors

## Usage

The command line tool, `grw`, can be used to easily read and write compressed resources by uri.  By default `grw`, reads from stdin and outputs to stdout.  The `--input-compression` and `--output-compression` flags are optional.

```shell
grw [--input-compression INPUT_COMPRESSION] [--output-compression OUTPUT_COMPRESSION] [flags]
```

To read from a resource provided by a URI use the first positional argument.

```shell
grw [--input-compression INPUT_COMPRESSION] [--output-compression OUTPUT_COMPRESSION] [flags] [-|stdin|INPUT_URI]
```

To write to a resource located by a URI, use the second positional argument.


```shell
grw [--input-compression INPUT_COMPRESSION] [--output-compression OUTPUT_COMPRESSION] [flags] [-|stdin|INPUT_URI] [-|stdout|OUTPUT_URI]
```


For more information use the help flag.

```shell
grw --help
```


## Algorithms

The following compression algorithms are supported.  Pull requests to support other compression algorithms are welcome!

| Algorithm | Read |  Write | Stream | Description |
| ---- | ------ | ------ | ------ | ------ |
| bzip2 | ✓ | - | ✓ | [bzip2](https://en.wikipedia.org/wiki/Bzip2) |
| flate | ✓ | ✓ | ✓ | [DEFLATE Compressed Data Format](https://tools.ietf.org/html/rfc1951) |
| gzip | ✓ | ✓ | ✓ | [gzip](https://en.wikipedia.org/wiki/Gzip) |
| snappy | ✓ | ✓ | ✓ | [snappy](https://github.com/google/snappy) |
| zip | ✓ | - | - | [zip](https://en.wikipedia.org/wiki/Zip_%28file_format%29) |
| zlib | ✓ | ✓ | ✓ | [zlib](https://en.wikipedia.org/wiki/Zlib) |


## Platforms

The following platforms are supported.  Pull requests to support other platforms are welcome!

| GOOS | 386 | amd64 | arm | arm64 |
| ---- | --- | ----- | --- | ----- |
| darwin | - | ✓ | - | - |
| freebsd | ✓ | ✓ | ✓ | - |
| linux | ✓ | ✓ | ✓ | ✓ |
| openbsd | ✓ | ✓ | - | - |
| solaris | - | ✓ | - | - |
| windows | ✓ | ✓ | - | - |

## Releases

Find releases for the supported platforms at [https://github.com/spatialcurrent/go-reader-writer/releases](https://github.com/spatialcurrent/go-reader-writer/releases).  See the **Building** section below to build for another platform from source.

# Examples

To download a file over https and write to stdout.

```shell
grw https://github.com/spatialcurrent/go-reader-writer/releases/download/0.0.1/grw.h -
```

To download a file from AWS S3, compress as gzip, and save locally.

```shell
grw --output-compression gzip s3://path/to/file /local/file
```

## Building

Use `make build_cli` to build executables for Linux and Windows.

**Changing Destination**

The default destination for build artifacts is `grw/bin`, but you can change the destination with an environment variable.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `DEST=/usr/local/go/bin make build_cli`

## Testing

To run CLI testes use `make test_cli`, which uses [shUnit2](https://github.com/kward/shunit2).  If you recive a `shunit2:FATAL Please declare TMPDIR with path on partition with exec permission.` error, you can modify the `TMPDIR` environment variable in line or with `export TMPDIR=<YOUR TEMP DIRECTORY HERE>`. For example:

```
TMPDIR="/usr/local/tmp" make test_cli
```

## Troubleshooting

### no such file or directory

#### Example

```text
error opening resource at uri %q: error opening file for writing at path %q: open %s: no such file or directory
```

#### Solution

This error typically occurs when a parent directory of an output file does not exist.  Use the `--output-mkdirs` command line flag to allow grw to create parent directories for output files as needed.
