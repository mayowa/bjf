// Tool to generate shortend urls based on ids
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xor-gate/go-bjf"
)

var mode  = flag.String("mode",  "encode", "Action encode/decode (default: encode)")
var base  = flag.String("base",  "Base64", "Base of alphabet (use: Base36, Base61, Base64)")
var token = flag.String("token", "125", "Token to encode/decode")

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var b bjf.Base = bjf.Base64

	flag.Parse()

	switch *base {
		case "Base36":
			b = bjf.Base64
		case "Base61":
			b = bjf.Base61
		case "Base64":
			b = bjf.Base64
	}

	bjf.Config(b)

	if *mode == "encode" {
		r := bjf.Encode(*token)
		fmt.Printf("%s -> %s\n", *token, r)
	} else if *mode == "decode" {
		r := bjf.Decode(*token)
		fmt.Printf("%s -> %d\n", *token, r)
	} else {
		Usage()
	}
}
