package main

import (
	"errors"
	"fmt"
)

func main() {
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Println(total)

	result, err := calculate(10, 2, "+")
	if err != nil {
		// return nil, err
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func sum(scores []int) int {
	var temp_score int
	for i := 0; i < len(scores); i++ {
		temp_score += scores[i]
	}
	return temp_score
}

func calculate(bil1 int, bil2 int, opr string) (int, error) {
	var bil int
	var errorResult error
	switch opr {
	case "+":
		bil = bil1 + bil2
	case "-":
		bil = bil1 - bil2
	case "*":
		bil = bil1 * bil2
	case "/":
		bil = bil1 / bil2
	default:
		errorResult = errors.New("Unknown operation")
	}
	return bil, errorResult
}
