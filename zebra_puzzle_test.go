package main

import "testing"

func TestSolvePuzzle(t *testing.T) {
	expected := puzzleSolution{DrinksWater: "norwegian", OwnsZebra: "japanese"}
	solution := SolveZebraPuzzle()
	if expected != solution {
		t.Fatalf("FAILED:\nExpected: %#v\nActual: %#v",
			expected, solution)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveZebraPuzzle()
	}
}
