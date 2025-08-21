package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
	exemplo07()
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// video 113
func exemplo01() {
	group := ColorGroup{1, "Reds", []string{"Crismon", "Red", "Ruby", "Maroon"}}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println(b)
	fmt.Println()
}

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

// video 114 marshall - Letras maiusculas em structs em go importam.
// marshal colocar no formato de json um dado
func exemplo02() {
	james := pessoa{"James", "Bond", 40, "Agente secreto", 1000000.50}
	darth := pessoa{"Anakin", "Skywalker", 50, "Sith", 7777777.87}

	j, err := json.Marshal(james)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	os.Stdout.Write(j)
	d, err1 := json.Marshal(darth)
	if err1 != nil {
		fmt.Println("Error: ", err1)
	}
	fmt.Println()

	os.Stdout.Write(d)
	fmt.Println(j, d)
}

type pessoaJson struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

// video 115 unmarshall
// unmarshall tirar do formato de json e transformar em um model.
func exemplo03() {
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente secreto","Contabancaria":1000000.5}`)

	var pessoa pessoa

	err := json.Unmarshal(sb, &pessoa)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%+v", pessoa)
}

// encode ja passa pro formato de json
func exemplo04() {
	darth := pessoa{"Anakin", "Skywalker", 50, "Sith", 7777777.87}
	fmt.Println()
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(darth)
}

// interface writer pacote sort video 116,117,118
func exemplo05() {
	ss := []string{"mofi", "total", "interior", "e", "capital"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)
	ss1 := []int{1, 4, 7, 9, 20, 2, 3, 4, 5, 7, 101, 101, 2031013, 2012, 12, 1201, 21, 2012, 03, 303013, 1301}
	fmt.Println(ss1)
	sort.Ints(ss1)
	fmt.Println(ss1)
}

//sort costumizado
//implementando a interface sort e conseguindo reutilizar sort atraves de uma implementação diferente

type carro struct {
	Nome     string
	Potencia int
	Consumo  int
}

// implementação por Potencia
type OrdenarPorPotencia []carro

func (x OrdenarPorPotencia) Len() int {
	return len(x)
}
func (x OrdenarPorPotencia) Less(i, j int) bool {
	return x[i].Potencia < x[j].Potencia
}

func (x OrdenarPorPotencia) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// implementação por Consumo
type OrdenarPorConsumo []carro

func (x OrdenarPorConsumo) Len() int {
	return len(x)
}
func (x OrdenarPorConsumo) Less(i, j int) bool {
	return x[i].Consumo < x[j].Consumo
}

func (x OrdenarPorConsumo) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// implementação por Lucro
type OrdenarporLucroProDonoDoPosto []carro

func (x OrdenarporLucroProDonoDoPosto) Len() int {
	return len(x)
}
func (x OrdenarporLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].Consumo < x[j].Consumo
}

func (x OrdenarporLucroProDonoDoPosto) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func Len() int {
	return 0
}
func exemplo06() {
	carros := []carro{
		carro{"Palio", 400, 4},
		carro{"Chevet", 200, 1},
		carro{"Mofi", 300, 3},
		carro{"Tatror", 560, 1},
	}

	fmt.Println("Inicial:\n", carros)
	sort.Sort(OrdenarPorPotencia(carros))
	fmt.Println("Potência:\n", carros)

	sort.Sort(OrdenarPorConsumo(carros))
	fmt.Println("Consumo:\n", carros)

	sort.Sort(OrdenarporLucroProDonoDoPosto(carros))
	fmt.Println("Lucro:\n", carros)
}

// usando bcrypt
func exemplo07() {
	senha := "22Outubro1991"
	senhaerrada := "21Outubro1991"
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	err = bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada))

	if err != nil {
		fmt.Println("Senha errada....")
	} else {
		fmt.Println("A senha está Correta...")
	}
	err = bcrypt.CompareHashAndPassword(sb, []byte(senha))

	if err != nil {
		fmt.Println("Senha errada....")
	} else {
		fmt.Println("A senha está Correta...")
	}
}
