package bjf

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	id := "125"
	r := Encode("125")
	fmt.Printf("%s -> %s\n", id, r)
	if r != "cb" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecode(t *testing.T) {
	id := "e9a"
	r := Decode("e9a")
	fmt.Printf("%s -> %d\n", id, r)
	if r != 19158 {
		t.Error("Unexpected decoded id", r)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("e9a")
	}
}
