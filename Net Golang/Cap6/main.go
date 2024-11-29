package main

import "fmt"

func main() {
	teste1()
	teste2()
	teste3()
	teste4()
	teste5()
	teste6()
	teste7()
}

func teste1() {
	for x := 100; x >= 0; x-- {

		if x == 5 {
			fmt.Println("x é igual a 5 saindo do loop")
			break
		}
		fmt.Println(x)
	}
}

func teste2() {
	x := 0
	for x < 10 {
		if x != 5 {
			fmt.Print(" ", x)

		} else {
			fmt.Print("x é igual a 5")
		}
		x++
	}
}

func teste3() {

	for x := 0; x < 20; x++ {
		if x%2 != 0 {
			//continue quebra a iteração atual
			continue
		}
		fmt.Println(x)

	}
}

func teste4() {
	for i := 1; i < 1000; i++ {
		fmt.Printf("%d, %#x, %#U, %v \t\n", i, i, i, string(i))
	}
}

func teste5() {
	y := 100

	if !(y > 80) {
		fmt.Println("y é maior que 80, o valor de y é ", y)
	} else if y > 80 {
		fmt.Println("y é maior que 80, o valor de y é ", y)
	}
}

func teste6() {

	z := 10
	switch {
	case z < 5:
		fmt.Println("z e menor que cinco")
		fallthrough
	case z == 5:
		fmt.Println("z e igual a cinco")

	case z > 5:
		fmt.Println("z e maior que cinco")
		fallthrough
	default:
		fmt.Println("z louco")
	}
}

func teste7() {

	x := 3
	if x == 2 || x == 3 {
		fmt.Println("x é igual a 2 ou 3")
	}

}
