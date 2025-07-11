package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
	exemplo07()
}

// video 87 e 88 funções em go
func exemplo01(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func exemplo02() {

	// si := []int{10, 10, 1, 2, 3, 5}
	// total := exemplo01(si...)
	// fmt.Println(total)
	total1 := exemplo01()

	fmt.Println(total1)
}

// video 89
func exemplo03() {
	defer fmt.Println("teste 1")
	fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")

}
func EscrevendoArquivo() {
	arquivo, err := os.Create("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}

	defer arquivo.Close()

	conteudo := `Linha 1: Olá Mundo Linha 2: Aprendendo Go Linha 3: Testando defer`
	_, err = arquivo.WriteString(conteudo)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo", err)
	}
}

func LerArquivo() {
	arquivo, err := os.Open("exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println(linha)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo")
	}
}
func exemplo04() {
	EscrevendoArquivo()
	LerArquivo()
}
func exemplo05() {

}
func exemplo06() {

}
func exemplo07() {

}
