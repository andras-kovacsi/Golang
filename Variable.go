package main

import "fmt"

func main() {

	var x int = 6	/* Declararea tipului variabilei si asocierea unei valori */
	var y int		/* Declararea tipului variabilei */
	y = 6			/* Asocierea unei valori */
	z := 7			/* Asocierea unei valori ce implicit recunoaste si tipul variabilei */
	
	var adun int = x+y
	var scad int = x-y
	var inm int = x*y
	var imp int = x/y
	var negX int = -x
	
	if x > y {
		fmt.Println("X este mai mare ca Y")
	} else if x < y {
		fmt.Println("X este mai mic ca Y")
	} else {
			fmt.Println("X este egal cu Y")
	}
	
	fmt.Println(x)
	fmt.Println(negX)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(adun)
	fmt.Println(scad)
	fmt.Println(inm)
	fmt.Println(imp)

}