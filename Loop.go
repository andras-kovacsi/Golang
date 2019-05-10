package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
		for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
	
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
}