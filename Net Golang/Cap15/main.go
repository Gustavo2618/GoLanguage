package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
}

// video 111 exercicios de ponteiro
func exemplo01() {
	x := 0
	fmt.Println(x)
	fmt.Println(&x)
}

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Hermes"
	p.sobrenome = "Luiz"
}

// video 112
func exemplo02() {
	gustavo := pessoa{"Gustavo", "Monteiro", 33}

	mudeMe(&gustavo)
	fmt.Println(gustavo.nome)
	fmt.Println(gustavo.sobrenome)
}
