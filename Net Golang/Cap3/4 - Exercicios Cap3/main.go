package main

import "fmt"

type intmofi int

var x intmofi

func main() {
	fmt.Printf(" %d, %T\n ", x, x)
	x := 42
	fmt.Println(x)

}
