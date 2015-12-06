package bjf

import (
	"fmt"
	"testing"
)

func TestEncodeDefaultBase62Zero(t *testing.T) {
	id := "0"
	r := Encode(id)
	fmt.Printf("%s -> %s\n", id, r)
	if r != "a" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestEncodeDefaultBase62(t *testing.T) {
	id := "125"
	r := Encode(id)
	fmt.Printf("%s -> %s\n", id, r)
	if r != "cb" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeDefaultBase62(t *testing.T) {
	id := "e9a"
	r := Decode(id)
	fmt.Printf("%s -> %d\n", id, r)
	if r != 19158 {
		t.Error("Unexpected decoded id", r)
	}
}

func TestEncodeBase59Zero(t *testing.T) {
	id := "0"
	Config(Base59)
	r := Encode(id)
	fmt.Printf("%s -> %s\n", id, r)
	if r != "a" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestEncodeBase59(t *testing.T) {
	id := "125"
	Config(Base59)
	r := Encode(id)
	fmt.Printf("%s -> %s\n", id, r)
	if r != "ch" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeBase59(t *testing.T) {
	id := "e9a"
	Config(Base59)
	r := Decode(id)
	fmt.Printf("%s -> %d\n", id, r)
	if r != 17346 {
		t.Error("Unexpected decoded id", r)
	}
}

func TestEncodeBase36Zero(t *testing.T) {
	id := "0"
	Config(Base59)
	r := Encode(id)
	fmt.Printf("%s -> %s\n", id, r)
	if r != "a" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestEncodeBase36(t *testing.T) {
	id := "125"
	Config(Base59)
	r := Encode("125")
	fmt.Printf("%s -> %s\n", id, r)
	if r != "ch" {
		t.Error("Unexpected encoded id", r)
	}
}

func TestDecodeBase36(t *testing.T) {
	id := "e9a"
	Config(Base59)
	r := Decode("e9a")
	fmt.Printf("%s -> %d\n", id, r)
	if r != 17346 {
		t.Error("Unexpected decoded id", r)
	}
}

func BenchmarkEncodeDefaultBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecodeBase62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("e9a")
	}
}

func BenchmarkEncodeBase59(b *testing.B) {
	Config(Base59)
	for i := 0; i < b.N; i++ {
		Encode("125")
	}
}

func BenchmarkDecodeBase59(b *testing.B) {
	Config(Base59)
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
