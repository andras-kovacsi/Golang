// https://www.quora.com/How-do-you-represent-mathematical-equations-in-golang
// 20.04.2019!!!

//	ax^2 + bx + c = 0
//	x^2 + 2x + 3 = 0
//	x1 = [-b + sqrt (b^2-4ac)] / 2
//	x1 = [-b - sqrt (b^2-4ac)] / 2

package main

import (
    "fmt"
	"math"
)
 
func main() {
	var a float64 = 2
	var b float64 = -1
	var c float64 = -3
	
	var negB float64 = -b
	var doiA float64 = a * 2
	var Bpatrat float64 = b * b
	var patruAC float64 = 4 * a * c
	var discrim float64 = Bpatrat - patruAC
	
	fmt.Println("Valoarea lui negB este: ", negB)
	fmt.Println("Valoarea lui doiA este: ", doiA)
	fmt.Println("Valoarea lui bpatrat este: ", Bpatrat)
	fmt.Println("Valoarea lui patruAC este: ", patruAC)
	fmt.Println("Valoarea lui discrim este: ", discrim)

	sq := math.Sqrt(discrim)
	if discrim < 0 {
		fmt.Println("\nValoriile X1 si X2 sunt imaginabile!\n")
	} else {
		var x1 = (negB + sq) / doiA
		var x2 = (negB - sq) / doiA
	fmt.Print("Valoarea lui X1 este: ", x1)
	fmt.Print("\nValoarea lui X2 este: ", x2)
	}
}