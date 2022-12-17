# `dmtx` - CLI Datamatrix PNG Creator

Create Datamatrix PNG from plain text files. Written in [Go](https://go.dev/) 1.19.

## Usage

    Usage: dmtx [options] [file]

    Options:
    -h, --help            display this help
    -s, --size int        size in pixel (default 600)
    -o, --output string   line separated file output
    -V, --version         display version

## Examples

    $ dmtx.exe --size 1200 README.md
    $ dmtx.exe --size 1200 README.md --output test.png

## Build

    make build

## Requirements

Only needed for embedding a resource icon on Windows:

    go install github.com/akavel/rsrc@latest
    rsrc -ico go.ico
    make build

---
Copyright © 2022 Typomedia Foundation. All rights reserved.
