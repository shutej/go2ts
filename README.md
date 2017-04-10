# go2ts

This is a simple tool that generates Typescript types, empty value
initializers, and JSON marshalling code from Go types.  This eliminates some of
the drudgery of keeping your frontend code consistent with an API implemented
in Go.

## Installation

```sh
go get -u github.com/shutej/go2ts/...
godep go install github.com/shutej/go2ts/cmd/go2ts
```

## Usage

The tool looks for a list of packages and types to convert from a YAML file, by
default looking for `go2ts.yml` in the current working directory.

You may consider making a synthetic package so `go generate` can compile your
assets for you:

```go
//go:generate go2ts --yml go2ts.yml --out gen
```
