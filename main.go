package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrInputZeroToCollatz = errors.New("The number must be unequal zero")
var ErrInvalidInput = errors.New("The given input is not valid")

func areIntSlicesEqual(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func calcolateCollatzStep(input int16) (int16, error) {
	// TODO: Write a test for the n == 0 check
	if input == 0 {
		err := fmt.Errorf("calcolateCollatzStep: %w", ErrInputZeroToCollatz)
		return 0, err
	}

	if input%2 == 0 {
		return (input / 2), nil
	}

	if input%2 == 1 {
		return (3*input + 1), nil
	}

	err := fmt.Errorf("calcolateCollatzStep: %w", ErrInvalidInput)
	return 0, err
}

func main() {
	var collatz_sequence []int16
	collatz_sequence = append(collatz_sequence, 121)
	for {
		currentNumber, err := calcolateCollatzStep(collatz_sequence[len(collatz_sequence)-1])
		if err != nil {
			switch {
			case errors.Is(err, ErrInputZeroToCollatz):
				fmt.Println(err)
			}

			log.Fatal()
		}

		collatz_sequence = append(collatz_sequence, currentNumber)

		if len(collatz_sequence) <= 2 {
			continue
		}

		if areIntSlicesEqual(collatz_sequence[len(collatz_sequence)-3:], []int16{4, 2, 1}) {
			break
		}
	}

	fmt.Println(collatz_sequence)
}
