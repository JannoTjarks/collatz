package collatz

import (
	"testing"
)

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
