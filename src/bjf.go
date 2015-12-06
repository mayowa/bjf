// Tool to generate shortend urls based on ids
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xor-gate/go-bijective"
)

var mode  = flag.String("mode",  "encode", "Action encode/decode (default: encode)")
var token = flag.String("token", "125", "Token to encode/decode")

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

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
