package main

import "fmt"

const x int = 31337
const y = 30

const (
	_ = 2024 + iota
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println(x)
	fmt.Println(y)

	//cap5 exercicio 4
	fmt.Printf("%d, %b, %#x\n", x, x, x)
	fmt.Printf("%d, %b, %#x\n", y, y, y)
	k := x << 8
	z := y >> 8

	fmt.Printf("%d, %b, %#x\n", k, k, k)
	fmt.Printf("%d, %b, %#x\n", z, z, z)
	//cap5 exercicio 5
	teste := `mofi ! malvina  e
	 tuf nunca aaaa				teste
	 se uniarao`
	fmt.Println(teste)
	//cap5 exercicio 6

	fmt.Println(b, c, d, e, f)

	//cap5 exercicio 7

}
