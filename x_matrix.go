package main

import (
	"errors"
	"fmt"
)

func main() {
	X := [][]float32{
		[]float32{1, 2, 3},
		[]float32{4, 5, 6},
	}

	w := [][]float32{
		[]float32{5, 2, 7},
		[]float32{5, 8, 3},
	}

	out, _ := multiply(X, transpose(w))
	fmt.Println(out)
}

func transpose(x [][]float32) [][]float32 {
	out := make([][]float32, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func multiply(x, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication.")
	}

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]float32, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out, nil
}