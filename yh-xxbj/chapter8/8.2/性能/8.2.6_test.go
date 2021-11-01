package main

import (
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test()
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestBlock()
	}
}
