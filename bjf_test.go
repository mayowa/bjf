package bjf

import (
	"fmt"
	"testing"
)

func TestEncodeDefaultBase64(t *testing.T) {
	id := "125"
	r := Encode("125")
	fmt.Printf("%s -> %s\n", id, r)
	if r != "cb" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeDefaultBase64(t *testing.T) {
	id := "e9a"
	r := Decode("e9a")
	fmt.Printf("%s -> %d\n", id, r)
	if r != 19158 {
		t.Error("Unexpected decoded id", r)
	}
}

func TestEncodeBase61(t *testing.T) {
	id := "125"
	Config(Base61)
	r := Encode("125")
	fmt.Printf("%s -> %s\n", id, r)
	if r != "ch" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeBase61(t *testing.T) {
	id := "e9a"
	Config(Base61)
	r := Decode("e9a")
	fmt.Printf("%s -> %d\n", id, r)
	if r != 17346 {
		t.Error("Unexpected decoded id", r)
	}
}

func TestEncodeBase36(t *testing.T) {
	id := "125"
	Config(Base61)
	r := Encode("125")
	fmt.Printf("%s -> %s\n", id, r)
	if r != "ch" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeBase36(t *testing.T) {
	id := "e9a"
	Config(Base61)
	r := Decode("e9a")
	fmt.Printf("%s -> %d\n", id, r)
	if r != 17346 {
		t.Error("Unexpected decoded id", r)
	}
}

func BenchmarkEncodeDefaultBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecodeBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("e9a")
	}
}

func BenchmarkEncodeBase61(b *testing.B) {
	Config(Base61)
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecodeBase61(b *testing.B) {
	Config(Base61)
	for i := 0; i < b.N; i++ {
		Decode("e9a")
	}
}

func BenchmarkEncodeBase36(b *testing.B) {
	Config(Base36)
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecodeBase36(b *testing.B) {
	Config(Base36)
	for i := 0; i < b.N; i++ {
		Decode("e9a")
	}
}
