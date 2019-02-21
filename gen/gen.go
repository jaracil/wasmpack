//go:generate go run gen.go

package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	base64js, err := ioutil.ReadFile("gen/base64.js")
	if err != nil {
		panic(err)
	}
	base64jsB64 := base64.StdEncoding.EncodeToString(base64js)

	inflatejs, err := ioutil.ReadFile("gen/inflate.js")
	if err != nil {
		panic(err)
	}
	inflatejsB64 := base64.StdEncoding.EncodeToString(inflatejs)

	bootstrap, err := ioutil.ReadFile("gen/wasm_exec.js")
	if err != nil {
		panic(err)
	}
	bootstrapB64 := base64.StdEncoding.EncodeToString(bootstrap)

	goTemplate, err := ioutil.ReadFile("template/wasmpack.go")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("wasmpack.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	packageTemplate := template.Must(template.New("").Parse(string(goTemplate)))

	packageTemplate.Execute(f, struct {
		Base64Code    string
		InflateCode   string
		BootStrapCode string
	}{
		Base64Code:    base64jsB64,
		InflateCode:   inflatejsB64,
		BootStrapCode: bootstrapB64,
	})

}
