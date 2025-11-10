package main

import "testing"

func Sum(a, b int) int {
	return a + b
}

func Multiple(a, b int) int {
	return a * b * 5000
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 12)
	}
}

func BenchmarkMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiple(62, 12)
	}
}