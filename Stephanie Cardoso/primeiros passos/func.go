package main

import "fmt"

func somar(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("the numbers must be positive")
	}

	return a + b, nil
}

func subtrair(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("the numbers must be positive")
	}

	return a - b, nil
}
