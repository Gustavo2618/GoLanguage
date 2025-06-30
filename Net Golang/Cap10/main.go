package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
	exemplo07()
	exemplo08()
}

//video 79
//introdução a structs
func exemplo01() {
	type cliente struct {
		nome      string
		sobrenome string
		fumante   bool
	}
	cliente01 := cliente{
		nome:      "Maria",
		sobrenome: "Elza",
		fumante:   false,
	}
	cliente02 := cliente{"Luiz", "Gustavo", false}

	fmt.Println(cliente01)
	fmt.Println(cliente02)
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func (P Produto) Exibir() {
	fmt.Println("Nome:", P.Nome)
	fmt.Println("Preco:", P.Preco)
	fmt.Println("Quantidade:", P.Quantidade)
}
func (P Produto) AtualizarEstoque(novoquant int) {
	P.Quantidade = novoquant
	fmt.Println("Valor de estoque atualizado para: ", novoquant)
}
func (P Produto) CalcularValor() float64 {
	return P.Preco * float64(P.Quantidade)

}
func exemplo02() {
	p := Produto{"Mouse", 500.00, 10}
	p.Exibir()
	p.AtualizarEstoque(20)
	fmt.Println("Valor total em estoque:", p.CalcularValor())
}

//video 80
func exemplo03() {

}
func exemplo04() {

}
func exemplo05() {

}
func exemplo06() {

}
func exemplo07() {

}
func exemplo08() {

}
