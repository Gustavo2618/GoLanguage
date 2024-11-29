package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprint(x, " ", y, " ", z)

	fmt.Println(s)
}
