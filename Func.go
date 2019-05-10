package main

import (
	"fmt"
	"errors"
	"math"
	)

func main() {
	result_1 := sum(3, 4)
	fmt.Println(result_1)
	
	result_2, err := sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result_2)
	}
}

func sum(x int, y int) int {
	return x+y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Nedefinit pentru valori negative")
	}
	return math.Sqrt(x), nil
}