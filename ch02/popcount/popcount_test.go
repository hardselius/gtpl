package popcount

import "testing"

func benchmark(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmark(b, PopCount)
}

func BenchmarkPopCountLoop(b *testing.B) {
	benchmark(b, PopCountLoop)
}

func BenchmarkPopCountByShift(b *testing.B) {
	benchmark(b, PopCountByShift)
}

func BenchmarkPopCountByClearing(b *testing.B) {
	benchmark(b, PopCountByClearing)
}