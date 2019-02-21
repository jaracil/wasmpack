package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	base64js, err := ioutil.ReadFile("base64.js")
	if err != nil {
		panic(err)
	}
	base64jsB64 := base64.StdEncoding.EncodeToString(base64js)

	inflatejs, err := ioutil.ReadFile("inflate.js")
	if err != nil {
		panic(err)
	}
	inflatejsB64 := base64.StdEncoding.EncodeToString(inflatejs)

	bootstrap, err := ioutil.ReadFile("wasm_exec.js")
	if err != nil {
		panic(err)
	}
	bootstrapB64 := base64.StdEncoding.EncodeToString(bootstrap)

	f, err := os.Create("embedded.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(
		`package main

var base64Code = "%s"

var inflateCode = "%s"

var bootstrapCode = "%s"
`, base64jsB64, inflatejsB64, bootstrapB64))

	if err != nil {
		panic(err)
	}

}
