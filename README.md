# wasmpack
Go wasm to standalone Javascript compressor.
Generates a JS file containing the compressed wasm file and all Go bootstrap JS code.

## Why?
We needed to distribute a Go wasm program in a single JS file.

## Install
go get github.com/jaracil/wasmpack

## Usage
wasmpack -i wasm_file > javascript_file

