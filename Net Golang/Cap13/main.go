package main

import (
	"fmt"
	"math"
)

// cap13 exercicios do cap 12
func main() {

	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
	exemplo07()
	exemplo08()
	exemplo09()
	exemplo10()

}

func mostranum(x int) int {
	return x
}

func doisretornos(x int, coisa string) (int, string) {
	return x, coisa

}

// exercicio 1
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

// exercicio 2 video 99 slices com entrelaçamento
func exemplo02() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 80}
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 80, 999}
	fmt.Println(soma(slice...))
	fmt.Println(soma1(slice1))
}

// exercicio 3 video 100
func exemplo03() {
	fmt.Println("Teste 2")
	defer fmt.Println("teste 0")
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) mostrarPessoa() {
	fmt.Println("Meu nome é:", p.sobrenome, p.nome, "e minha idade é: ", p.idade, "anos")

}

// exercicio 4 video 101
func exemplo04() {
	gustavo := pessoa{
		"Gustavo", "Luiz", 33,
	}

	gustavo.mostrarPessoa()
}

type circulo struct {
	raio float64
}

type quadrado struct {
	lado1 float64
}

func (c circulo) calculaArea() {
	fmt.Println("Area circulo: ", c.raio*2*math.Pi)
}
func (q quadrado) calculaArea() {
	fmt.Println("Area quadrado: ", q.lado1*q.lado1)
}

type info interface {
	calculaArea()
}

func medida(i info) {
	i.calculaArea()
}

// exercicio 5 video 102
func exemplo05() {
	c := circulo{
		10,
	}

	q := quadrado{
		7.2,
	}

	//fmt.Println("Area circulo: ", c.calculaArea())
	//fmt.Println("Area quadrado: ", q.calculaArea())
	medida(c)
	medida(q)
}

// exercicio 6 video 103
func exemplo06() {
	lado1 := 10.0

	func(lado1 float64) {
		fmt.Println("Area quadrado:", lado1*lado1)
		fmt.Println("Se fosse um circulo", lado1*2*math.Pi)
	}(lado1)
}

// exercicio 7 video 104
func exemplo07() {
	teste := func() {
		fmt.Println("Gustavo namora Patricia porque ela é bonita!")
	}
	teste()
}

func testeSoma() func(...int) int {
	return func(i ...int) int {
		total := 0
		for _, v := range i {
			total += v
		}
		return total
	}
}

// exercicio 8 video 105
func exemplo08() {
	x := testeSoma()
	t := x(1, 2, 3, 4, 5, 6, 7, 8, 88, 10)
	fmt.Println(t)

	fmt.Println(testeSoma()(1, 2, 3, 4, 5, 6, 7, 8, 88, 10, 99))
}

func argumento() {
	fmt.Println("Olha eu aqui!")
}
func recebeArgumento(x func()) {
	fmt.Println("Prestenção:")
	x()
}

// exercicio 9 video 106
func exemplo09() {
	recebeArgumento(argumento)
}

func clousure() func() int {
	x := 0
	return func() int {
		x--
		return x
	}
}

// exercicio 10 video 107 e 108
func exemplo10() {
	teste1 := clousure()
	fmt.Println(teste1())
	fmt.Println(teste1())
	fmt.Println(teste1())
	teste2 := clousure()
	fmt.Println(teste2())
	fmt.Println(teste2())
	fmt.Println(teste2())

}
