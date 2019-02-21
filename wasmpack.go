//go:generate go run gen/gen.go

package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	InputFile string `short:"i" long:"input" description:"WASM input file"`
}

func genCode(s string) string {
	bin, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(bin)
}

func main() {
	flags.Parse(&opts)
	if opts.InputFile == "" {
		fmt.Fprintf(os.Stderr, "Input file not specified\n")
		return
	}
	f, err := os.Open(opts.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %#v\n", err)
		return
	}
	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %#v\n", err)
		return
	}
	buf := &bytes.Buffer{}
	def, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating deflate object: %#v\n", err)
		return
	}
	n, err := def.Write(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compressing data: %#v\n", err)
		return
	}
	if len(input) != n {
		fmt.Fprintf(os.Stderr, "buffer size missmatch")
		return
	}
	err = def.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing data: %#v\n", err)
		return
	}
	def.Close()
	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Fprintf(os.Stdout,
		`/* Compressed with wasmpack [https://github.com/jaracil/wasmpack.git] */

%s
%s
%s
var go = new Go();
window.onload = function(){
	WebAssembly.instantiate(new Zlib.RawInflate(base64js.toByteArray("%s")).decompress() , go.importObject).then(function(results){
		go.run(results.instance);
	});
};
`,
		genCode(base64Code),
		genCode(inflateCode),
		genCode(bootstrapCode),
		b64str,
	)

}
