package main

import "testing"

func TestCalculator(t *testing.T) {
	t.Run(".Add", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run(".Subtract", func(t *testing.T) {
		sum := Subtract(4, 2)
		expected := 2

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
