# wasmpack
Go wasm to standalone Javascript file.
Generates a JS file containing the compressed wasm file and all Go bootstrap JS code to execute in browser and node.js environments.

## Why?
We needed to distribute a Go wasm program in a single JS file.

## Install
```
go get -u github.com/jaracil/wasmpack
```

If you modify base64.js, inflate.js or wasm_exec.js then execute:
```
go generate
go install
```

## Usage
```
wasmpack -i file.wasm > file.js
node file.js
```

