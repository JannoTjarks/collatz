package main

import (
	"testing"
)

func TestCalcolateCollatzStepEvenInput(t *testing.T) {
	got, err := calcolateCollatzStep(2)
	var want int16 = 1
	if want != got || err != nil {
		t.Fatalf(`calcolateCollatzStep(2) = %d, %v, want %d, nil`, got, err, want)
	}
}

func TestCalcolateCollatzStepOddInput(t *testing.T) {
	got, err := calcolateCollatzStep(7)
	var want int16 = 22
	if want != got || err != nil {
		t.Fatalf(`calcolateCollatzStep(2) = %d, %v, want %d, nil`, got, err, want)
	}
}

func TestCalcolateCollatzStepZeroInput(t *testing.T) {
	got, err := calcolateCollatzStep(0)
	var want int16 = 0
	if want != got || err == nil {
		t.Fatalf(`calcolateCollatzStep(2) = %d, %v, want %d, error`, got, err, want)
	}
}

func TestAreIntSlicesEqual(t *testing.T) {
	got := areIntSlicesEqual([]int16{0, 1, 2}, []int16{0, 1, 2})
	want := true
	if want != got {
		t.Fatalf(`areIntSlicesEqual([]int16{0, 1, 2}, []int16{0, 1, 2}) = %t, want %t`, got, want)
	}
}

func TestAreIntSlicesEqualDifferentValues(t *testing.T) {
	got := areIntSlicesEqual([]int16{0, 5, 6}, []int16{0, 1, 2})
	want := false
	if want != got {
		t.Fatalf(`areIntSlicesEqual([]int16{0, 5, 6}, []int16{0, 1, 2}) = %t, want %t`, got, want)
	}
}

func TestAreIntSlicesEqualDifferentOrder(t *testing.T) {
	got := areIntSlicesEqual([]int16{2, 1, 0}, []int16{0, 1, 2})
	want := false
	if want != got {
		t.Fatalf(`areIntSlicesEqual([]int16{0, 5, 6}, []int16{0, 1, 2}) = %t, want %t`, got, want)
	}
}
