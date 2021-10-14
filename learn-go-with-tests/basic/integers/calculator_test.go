package main

import "testing"

func TestCalculator(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run(".Add", func(t *testing.T) {
		actual := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, actual, expected)
	})

	t.Run(".Subtract", func(t *testing.T) {
		actual := Subtract(4, 2)
		expected := 2

		assertCorrectMessage(t, actual, expected)
	})
}
