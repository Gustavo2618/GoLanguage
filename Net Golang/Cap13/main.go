package main

import "fmt"

//cap13 exercicios do cap 12
func main() {

	exemplo01()
	exemplo02()
	exemplo03()

}

func mostranum(x int) int {
	return x
}

func doisretornos(x int, coisa string) (int, string) {
	return x, coisa

}

//exercicio 1
func exemplo01() {
	fmt.Println(mostranum(10))
	fmt.Println(doisretornos(15, "Mofi"))
}

func soma(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func soma1(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

//exercicio 2 video 99 slices com entrelaçamento
func exemplo02() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 80}
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 80, 999}
	fmt.Println(soma(slice...))
	fmt.Println(soma1(slice1))
}

//exercicio 3 video 100
func exemplo03() {
	fmt.Println("Teste 2")
	defer fmt.Println("teste 0")
}
