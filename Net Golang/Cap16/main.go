package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
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
