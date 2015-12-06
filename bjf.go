// Package bjf provides base62 bijective function Encode/Decode for
// use in url shorters
// Based on http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener
package bjf

import (
	"math"
	"strings"
	"strconv"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(num string) string {
	n, _ := strconv.ParseUint(num, 10, 64)
	t := make([]byte, 0)

	/* Special case */
	if n == 0 {
		return string(alphabet[0])
	}

	/* Map */
	for n > 0 {
		r := n % uint64(len(alphabet))
		t = append(t, alphabet[r])
		n = n / uint64(len(alphabet))
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
		r += strings.Index(alphabet, string(token[i])) * int(math.Pow(float64(len(alphabet)), p))
		p--
	}

	return r
}
