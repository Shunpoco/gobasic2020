package popcount

import "testing"

var x uint64 = 12345678987654321

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountBitShift(x)
	}
}

func BenchmarkPopCountUnit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountUnit(x)
	}
}
