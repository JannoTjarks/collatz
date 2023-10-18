package collatz

import (
	"errors"
	"fmt"
	"log"
)

func GetCollatzSequence(startValue int16) ([]int16, error) {
	var collatz_sequence []int16
	collatz_sequence = append(collatz_sequence, startValue)
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

	return collatz_sequence, nil
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

func isConjectureFulfilled(sequence []int16) (bool, error) {
	if areIntSlicesEqual(sequence[len(sequence)-3:], []int16{4, 2, 1}) {
		return true, nil
	}

	return false, nil
}
