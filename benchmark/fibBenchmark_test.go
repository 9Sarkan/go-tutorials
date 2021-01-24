package benchmark

import "testing"

var result int

func benchmarkfib1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib1(n)
	}
	result = r
}
func benchmarkfib2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib2(n)
	}
	result = r
}
func benchmarkfib3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fib3(n)
	}
	result = r
}

func Benchmark30fib1(b *testing.B) {
	benchmarkfib1(b, 30)
}
func Benchmark30fib2(b *testing.B) {
	benchmarkfib2(b, 30)
}
func Benchmark30fib3(b *testing.B) {
	benchmarkfib3(b, 30)
}
