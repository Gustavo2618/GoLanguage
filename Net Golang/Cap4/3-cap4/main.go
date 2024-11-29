package main

import "fmt"

//constantes podem ser atribuidas a mais de um tipo como por exemplo int e float64 isso acontece pq o compilador so vai
// saber oque a constante é quando ela e utilizada.
const x = 10

const (
	k = 20
	w = 15.1
	r = true
)

var y int
var z float64

func main() {
	// y = x
	// fmt.Printf("%d , %T", y, y)
	z = x
	fmt.Printf("%f , %T\n", z, z)

	fmt.Println(k, w, r)
}
