package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()

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
//structs embutidos
type Pessoa struct {
	nome  string
	idade int
}
type Profissional struct {
	pessoa  Pessoa
	titulo  string
	salario int
}

func exemplo03() {
	gustavo := Pessoa{"Gustavo", 33}
	profissaoGustavo := Profissional{gustavo, "Bacharel em Computação", 2150}

	fmt.Println(profissaoGustavo)
}

type Funcionario struct {
	pessoa Pessoa
	cargo  string
}

func (f Funcionario) Exibir() {
	fmt.Println("Pessoa:", f.pessoa)
	fmt.Println("Cargo:", f.cargo)
}

type Gerente struct {
	funcionario  Funcionario
	departamento string
}

func (g Gerente) Exibir() {
	fmt.Println("Gerente:", g.funcionario)
	fmt.Println("Cargo:", g.departamento)
}

func exemplo04() {
	gustavo := Pessoa{nome: "Gustavo", idade: 33}
	funcGustavo := Funcionario{gustavo, "Desenvolvedor Junior"}
	gerenteGustavo := Gerente{funcGustavo, "Departamendo HP"}

	fmt.Println("Exibir Funcionario:", funcGustavo)
	fmt.Println("Exibir Gerente:", gerenteGustavo)

}
