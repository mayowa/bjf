// Tool to generate shortend urls based on ids
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xor-gate/go-bjf"
)

var mode  = flag.String("mode",  "encode", "Action encode/decode (default: encode)")
var base  = flag.String("base",  "Base62", "Base of alphabet (use: Base36, Base59, Base62)")
var token = flag.String("token", "125", "Token to encode/decode")

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var b bjf.Base = bjf.Base62

	flag.Parse()

	switch *base {
		case "Base36":
			b = bjf.Base62
		case "Base59":
			b = bjf.Base59
		case "Base62":
			b = bjf.Base62
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
