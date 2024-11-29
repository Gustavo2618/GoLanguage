package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) //zero value
	x = true
	fmt.Println(x) //valor atribuido

	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	e := "e" //ascii
	b := "é" //utf-8
	c := "香" //caractere oriental
	fmt.Printf("%v,%v,%v\n ", e, b, c)
	//conversao de letras para bytes
	d := []byte(e)
	g := []byte(b)
	f := []byte(c)
	fmt.Printf("%v,%v,%v\n ", d, g, f)

	//overflow
	var i uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

}
