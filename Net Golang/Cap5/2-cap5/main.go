package main

import "fmt"

func main() {

	x := 10
	for x = 0; x < 100; x++ {
		if x == 10 {
			fmt.Println("x == 10")
		}
		if x > 1 {
			fmt.Println("x > 1")
		}
		if x != 99 {
			fmt.Println("x != de 99")
		}
		if x >= 98 {
			fmt.Println("x > 98")
		}
		if x >= 97 {
			fmt.Println("x >= 97")
		}
		if x < 3 {
			fmt.Println("x < 3")
		}
		if x <= 2 {
			fmt.Println("x < 2")
		}
	}
}
