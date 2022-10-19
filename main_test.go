package main

import "testing"

func BenchmarkWithIf(b *testing.B) {
	a := make([]int, 100000000)
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		_ = withIf(a)
	}
}

func BenchmarkWithForElement(b *testing.B) {
	a := make([]int, 100000000)
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		_ = withForElement(a)
	}
}

func BenchmarkWithForIndex(b *testing.B) {
	a := make([]int, 100000000)
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		_ = withForIndex(a)
	}
}

func BenchmarkWithIfZero(b *testing.B) {
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		var a []int
		_ = withIf(a)
	}
}

func BenchmarkWithForElementZero(b *testing.B) {
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		var a []int
		_ = withForElement(a)
	}
}

func BenchmarkWithForIndexZero(b *testing.B) {
	b.N = 10000000000
	for i := 0; i < b.N; i++ {
		var a []int
		_ = withForIndex(a)
	}
}
