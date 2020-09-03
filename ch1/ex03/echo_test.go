package main

import "testing"

var inputs []string

func init() {
	for i := 0; i < 10000; i++ {
		inputs = append(inputs, "hoge")
	}
}

func BenchmarkSlowEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowEcho(inputs)
	}
}

func BenchmarkFastEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastEcho(inputs)
	}
}
