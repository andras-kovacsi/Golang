package main

import "fmt"

func main() {
	wert := make(map[string]int) /* [Tipul cheilor]tipul valorii */
	wert["triunghi"] = 3
	wert["patrat"] = 4
	wert["hexagon"] = 6
	
	delete(wert, "triunghi")
	
	fmt.Println(wert)	/* map[hexagon:6 patrat:4 triunghi:3] */
	fmt.Println(wert["patrat"])	/* 4 */
}