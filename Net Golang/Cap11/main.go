package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()

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

type veiculo struct {
	portas int
	cor    string
}
type sedan struct {
	veiculo    veiculo
	modeloLuxo bool
}
type caminhonete struct {
	veiculo           veiculo
	tracaoQuatroRodas bool
}

//video 85
func exemplo03() {
	veiculo1 := veiculo{4, "Vermelho"}
	veiculo2 := veiculo{4, "Branco"}

	sedan := sedan{veiculo1, true}
	caminhonete := caminhonete{veiculo2, true}

	fmt.Println("Meu sedan:", sedan)
	if sedan.modeloLuxo {
		fmt.Println("Meu sedan é um veículo de luxo!")
	}
	fmt.Print("Minha caminhonete:", caminhonete, "A cor da caminhonete é:", caminhonete.veiculo.cor, "\n")
	if caminhonete.tracaoQuatroRodas {
		fmt.Println("Minha caminhonete tem quatro Portas")
	}
}

//video86
func exemplo04() {
	anonimo := struct {
		mapa  map[string]string
		slice []int
	}{
		mapa: map[string]string{
			"teste":  "teste1",
			"teste2": "teste3",
		},
		slice: []int{1, 2, 3, 4, 5, 6},
	}

	fmt.Println(anonimo)
}
