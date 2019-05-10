package main

import (
	"fmt"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	html_inceput = `<!DOCTYPE HTML>
<html>
 <head>
  <style>.error{color:#FF0000;}</style>
  <title>Rezolvare Ecuatie Gradul 2</title>
 </head>
 <body>
  <h1>Rezolvare ecuatie de gradul 2</h1>
  <p>a<i>x</i><sup>2</sup> + b<i>x</i> + c</p>

  <form action="/" method="POST">
   <label for="a">a = </label><input type="text" name="a" size="2"> </br>
   <label for="b">b = </label><input type="text" name="b" size="2"> </br>
   <label for="c">c = </label><input type="text" name="c" size="2"> </br>
   <input type="submit" name="rezolva" value="Rezolva">
  </form>
`
	html_sfarsit = `
 </body>
</html>`

	solutieX1 = "x1="
	solutieX2 = "</br>x2="
	faraSolutie = "<p>Ecuatia nu are solutie</p>"
)

func main() {
	http.HandleFunc("/", rezolvareEcuatieGradDoi)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Nu s-a putut porni serverul", err)
	}
}


func rezolvareEcuatieGradDoi(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Fprint(writer, html_inceput)

	if request.Method == "POST" {

		a, _ := strconv.ParseFloat(request.FormValue("a"), 64);
		b, _ := strconv.ParseFloat(request.FormValue("b"), 64);
		c, _ := strconv.ParseFloat(request.FormValue("c"), 64);
		x1, x2 := radacini(a, b, c)

		if cmplx.IsNaN(x1) && cmplx.IsNaN(x2) {
			fmt.Fprint(writer, faraSolutie)
		}
		fmt.Fprint(writer, "Solutiie ecuatiei sunt:<br/>")
		fmt.Fprint(writer, solutieX1)
		fmt.Fprint(writer, x1)
		fmt.Fprint(writer, solutieX2)
		fmt.Fprint(writer, x2)
	}
	fmt.Fprint(writer, html_sfarsit)
}

func radacini(a float64, b float64, c float64) (complex128, complex128) {
	aa, bb, cc := complex(a, 0), complex(b, 0),
		complex(c, 0)
	root := cmplx.Sqrt(cmplx.Pow(bb, 2) - (4 * aa * cc))
	x1 := (-bb + root) / (2 * aa)
	x2 := (-bb - root) / (2 * aa)
	return x1, x2
}