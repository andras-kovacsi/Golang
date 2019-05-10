package main

import "fmt"

type persoana struct {
	nume string
	varsta int
}

func main() {
	p := persoana{nume: "Erik", varsta: 2}
	fmt.Println(p)
	fmt.Println(p.nume)
}