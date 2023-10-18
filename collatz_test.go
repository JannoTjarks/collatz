package collatz

import (
	"testing"
)

func TestCalcolateCollatzStepEvenInput(t *testing.T) {
	got, err := runCollatzOperation(2)
	var want int16 = 1
	if want != got || err != nil {
		t.Fatalf(`calcolateCollatzStep(2) = %d, %v, want %d, nil`, got, err, want)
	}
}

func TestCalcolateCollatzStepOddInput(t *testing.T) {
	got, err := runCollatzOperation(7)
	var want int16 = 22
	if want != got || err != nil {
		t.Fatalf(`calcolateCollatzStep(7) = %d, %v, want %d, nil`, got, err, want)
	}
}

func TestCalcolateCollatzStepZeroInput(t *testing.T) {
	got, err := runCollatzOperation(0)
	var want int16 = 0
	if err == nil || want != got {
		t.Fatalf(`calcolateCollatzStep(0) = %d, %v, want %d, error`, got, err, want)
	}
}

func TestIsConjectureFulfilledValid(t *testing.T) {
	got, _ := isConjectureFulfilled([]int16{4, 2, 1})
	var want bool = true
	if want != got {
		t.Fatalf(`isConjectureFulfilled([]int16{4,2,1}) = %t, want %t`, got, want)
	}
}

func TestIsConjectureFulfilledValid2(t *testing.T) {
	got, _ := isConjectureFulfilled([]int16{3, 4, 18, 30, 4, 2, 1})
	var want bool = true
	if want != got {
		t.Fatalf(`isConjectureFulfilled([]int16{4,2,1}) = %t, want %t`, got, want)
	}
}

func TestIsConjectureFulfilledInvalid(t *testing.T) {
	got, _ := isConjectureFulfilled([]int16{8, 3, 1})
	var want bool = false
	if want != got {
		t.Fatalf(`isConjectureFulfilled([]int16{4,2,1}) = %t, want %t`, got, want)
	}
}
