# CLI

- [Algorithms](#algorithms) - list of supported compression algorithms
- [Platforms](#platforms) - list of supported platforms
- [Releases](#releases) - where to find an executable
- [Examples](#examples)  - detailed usage exampels
- [Examples](#building) - how to build the CLI
- [Testing](#testing) - test the CLI
- [Troubleshooting](#Troubleshooting) - how to troubleshoot common errors

## Usage

The command line tool, `grw`, can be used to easily read and write compressed resources by uri.

### Algorithms

The following compression algorithms are supported.  Pull requests to support other algorithms are welcome!

| Algorithm | Read |  Write |
| ---- | ------ |  ------ |
| bzip2 | ✓ | - |
| flate | ✓ | ✓ |
| gzip | ✓ | ✓ |
| snappy | ✓ | ✓ |
| zip | ✓ | - |
| zlib | ✓ | ✓ |


### Platforms

The following platforms are supported.  Pull requests to support other platforms are welcome!

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

## Releases

**grw** is currently in **alpha**.  Find releases at [https://github.com/spatialcurrent/go-reader-writer/releases](https://github.com/spatialcurrent/go-reader-writer/releases).  See the **Building** section below to build from scratch.

**Darwin**

- `grw_darwin_amd64` - CLI for Darwin on amd64 (includes `macOS` and `iOS` platforms)

**Linux**

- `grw_linux_amd64` - CLI for Linux on amd64
- `grw_linux_amd64` - CLI for Linux on arm64

**Windows**

- `grw_windows_amd64.exe` - CLI for Windows on amd64

# Examples

To download a file over https and write to stdout.

```shell
grw https://github.com/spatialcurrent/go-reader-writer/releases/download/0.0.1/grw.h -
```

To download a file from AWS S3, compress as gzip, and save locally.

```shell
grw --output-compression gzip s3://path/to/file /local/file
```

# Building

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

