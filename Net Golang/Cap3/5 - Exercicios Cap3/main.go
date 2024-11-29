package main

import "fmt"

type intmofi int

var x intmofi
var y intmofi

func main() {
	fmt.Printf(" %d, %T\n ", x, x)
	x := 42
	fmt.Println(x)
	y := x
	fmt.Printf("variavel y: %d, %T\n ", y, y)
}