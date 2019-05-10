package main

import "fmt"

func main() {
	var a [5]int	/* Numar de elemente */
	a[2] = 7		/* Al treilea element va fi 7 [0, 1, 2, 3, 4]*/
	b := [5]int{2, 4, 6, 8, 10}
	
/*	Slice-uri, nu au un numar fix de elemente	*/

	c := []int{2, 3, 5, 7, 9}
	c = append(c, 11)
	
	/* Se va initializa cu 0 daca nu precizam nici o valoare*/
	fmt.Println(a)	/*[0 0 7 0 0]*/
	fmt.Println(b)	/*[2 4 6 8 10]*/
	fmt.Println(c)	/*[1 3 5 7 9 11]*/
}