// Package bjf provides base62 bijective function Encode/Decode for
// use in url shorters
// Based on http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener
package bjf

import (
	"math"
	"strings"
	"strconv"
)

type Base string
const (
	Base36 Base = "abcdefghijklmnopqrstuvwxyz0123456789"
	Base59 Base = "abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	Base62 Base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var alphabetConfig = Base62

func Config(b Base) {
	alphabetConfig = b
}

func Encode(num string) string {
	n, _ := strconv.ParseUint(num, 10, 64)
	t := make([]byte, 0)

	/* Special case */
	if n == 0 {
		return string(alphabetConfig[0])
	}

	/* Map */
	for n > 0 {
		r := n % uint64(len(alphabetConfig))
		t = append(t, alphabetConfig[r])
		n = n / uint64(len(alphabetConfig))
	}

	/* Reverse */
	for i, j := 0, len(t) - 1; i < j; i, j = i + 1, j - 1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}

func Decode(token string) int {
	r := int(0)
	p := float64(len(token)) - 1

	for i := 0; i < len(token); i++ {
		r += strings.Index(string(alphabetConfig), string(token[i])) * int(math.Pow(float64(len(alphabetConfig)), p))
		p--
	}

	return r
}
