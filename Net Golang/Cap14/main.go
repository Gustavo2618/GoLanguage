package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
}

// video 109 ponteiros
func exemplo01() {
	x := 10
	y := &x
	*y++
	fmt.Println(x, y)
	*y++
	defer fmt.Println(*y)
	fmt.Println(&x)

	//conteudo do endereço de x
	fmt.Println(*&x)

	defer fmt.Printf("%T, %T\n", x, y)

}

func recebevalor(x int) {
	x++
	fmt.Println("No valor:", x)
}

func recebeponteiro(x *int) {
	*x++
	fmt.Println("No Ponteiro", *x)
}

//video 110 utilidades de ponteiros
func exemplo02() {

	x := 0
	//criação de uma copia para a função
	recebevalor(x)
	fmt.Println("copia: ", x)

	recebeponteiro(&x)

	fmt.Println("depois de incrementar ponteiro:", x)
}
