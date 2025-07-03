package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
}

// video83
type Pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func exemplo01() {
	pessoa1 := Pessoa{"Luiz", "Gustavo", []string{"Morango", "Tapioca"}}
	pessoa2 := Pessoa{"Ana", "Patricia", []string{"Brigadeiro", "Ninho"}}

	fmt.Println(pessoa1.Nome, pessoa1.Sobrenome)
	for i := range pessoa1.Sabores {
		fmt.Println(pessoa1.Sabores[i])
	}
	fmt.Println(pessoa2.Nome, pessoa2.Sobrenome)
	for i := range pessoa2.Sabores {
		fmt.Println(pessoa2.Sabores[i])
	}
}

//video84 utilizando mapas no struct do ultimo exercicio
func exemplo02() {
	mapaPessoa := map[string]Pessoa{
		"Luiz":     Pessoa{"Gustavo", "Luiz", []string{"Morango", "Tapioca"}},
		"Patricia": Pessoa{"Ana", "Patricia", []string{"Brigadeiro", "Ninho"}},
	}
	for _, valor := range mapaPessoa {
		fmt.Println("Meu nome é:", valor.Nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.Sabores {
			fmt.Println("-", valor, "\n")
		}
	}

}
func exemplo03() {

}
func exemplo04() {

}
func exemplo05() {

}
func exemplo06() {

}
