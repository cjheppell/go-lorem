# go-lorem
Lorem Ipsum generator written in Golang

## Build

Run the following at the root of the repo:

```
go build
```

## Usage

Invoke the binary `./go-lorem` and pass it the `--size` flag with the requested size.

The format is as follows: `<size-as-integer>[B|K|M|G]` where B, K, M, G represent bytes, kilobytes, megabytes and gigabytes respectively

e.g:

`./go-lorem --size 1M` will generate ~1MB of Lorem Ipsum strings to stdout.

If you want to write this to a file, then redirect the stdout stream, e.g: `./go-lorem --size 1M > lorem.1m.txt`