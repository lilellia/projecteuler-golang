package main

import "testing"

// ===== Problem 001 ==============================================

func TestPE001(t *testing.T) {
	got := PE001(3, 5, 1000)
	want := 233168
	if got != want {
		t.Errorf("PE001(3, 5, 1000) = %v // want %v", got, want)
	}
}

func BenchmarkPE001(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE001(3, 5, 1000)
	}
}

// ===== Problem 002 ==============================================

func isEven(n int) bool { return n % 2 == 0 }

func TestPE002(t *testing.T) {
	got := PE002(4000000, isEven)
	want := 4613732
	if got != want {
		t.Errorf("PE002(4000000, isEven) = %v // want %v", got, want)
	}
}

func BenchmarkPE002(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE002(4000000, isEven)
	}
}

// ===== Problem 003 ==============================================

const n003 = 600851475143

func TestPE003(t *testing.T) {
	got := PE003(n003)
	want := 6857
	if got != want {
		t.Errorf("PE003(%d) = %v // want %v", n003, got, want)
	}
}

func BenchmarkPE003(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PE003(n003)
	}
}