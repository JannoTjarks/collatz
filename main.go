package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

// TODO: Create Tests for isConjectureFulfilled()
func isConjectureFulfilled(sequence []int16) (bool, error) {
	if areIntSlicesEqual(sequence[len(sequence)-3:], []int16{4, 2, 1}) {
		return true, nil
	}

	return false, nil
}

func runCollatzOperation(input int16) (int16, error) {
	if input == 0 {
		err := fmt.Errorf("runCollatzOperation: %w", ErrInputZeroToCollatz)
		return 0, err
	}

	if input%2 == 0 {
		return (input / 2), nil
	}

	if input%2 == 1 {
		return (3*input + 1), nil
	}

	err := fmt.Errorf("runCollatzOperation: %w", ErrInvalidInput)
	return 0, err
}

// TODO: Write Tests for the func parseCommandLineArgumentsToCollatzInput()
// TODO: Check if the collatzInput is given
// TODO: Catch a empty array as arguments
func parseCommandLineArgumentsToCollatzInput(arguments []string) (int16, error) {
	collatzInput, err := strconv.ParseInt(arguments[0], 10, 16)
	if err != nil {
		err := fmt.Errorf("parseCommandLineArgumentsToCollatzInput: %w", ErrInvalidInput)
		return 0, err
	}

	return int16(collatzInput), nil
}

func main() {
	collatzStartValue, _ := parseCommandLineArgumentsToCollatzInput(os.Args[1:])

	var collatz_sequence []int16
	collatz_sequence = append(collatz_sequence, collatzStartValue)
	for {
		currentNumber, err := runCollatzOperation(collatz_sequence[len(collatz_sequence)-1])
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

		fulfillment, _ := isConjectureFulfilled(collatz_sequence)
		if fulfillment == true {
			break
		}
	}

	fmt.Println(collatz_sequence)
}
