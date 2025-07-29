package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {

	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()
	exemplo07()
	exemplo08()
	exemplo09()
	exemplo10()
	exemplo11()
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

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type dentista struct {
	pessoa         pessoa
	especializacao string
	salario        float64
}
type arquiteto struct {
	pessoa           pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é :", x.pessoa.nome, "e ouve só: Bom dia!")
}
func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é :", x.pessoa.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
}

// polimorfismo e interfaces em go
func exemplo05() {
	mrpredio := dentista{
		pessoa:         pessoa{"mrDente", "da silva", 50},
		especializacao: "Cirurgia facial",
		salario:        8000.00,
	}
	mrdente := arquiteto{
		pessoa:           pessoa{"mrpredio", "15 andares", 38},
		tipodeconstrução: "Metro",
		tamanhodaloucura: "media",
	}
	mrdente.oibomdia()
	mrpredio.oibomdia()

	serhumano(mrdente)
	serhumano(mrpredio)

}

// funcoes anonimas
func exemplo06() {

	x := 387
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)
}

// funções como expressão
func exemplo07() {
	x := 387
	y := func(x int) int {
		return x * 873648
	}

	fmt.Println(x, "vezes 873648 é: ", y(x))
}

func retornaumafunção() func(int) int {

	return func(i int) int {
		return i * 10
	}
}

// função que retorna função
func exemplo08() {
	x := retornaumafunção()

	t := x(3)
	fmt.Println(t)
	fmt.Println(retornaumafunção()(3))
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 1 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

// callbacks
func exemplo09() {
	t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println("soma de somente pares: ", t)
	k := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println("soma de somente impares: ", k)
}
func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// clousure  = escopos?
func exemplo10() {
	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
func fatorial(x int64) *big.Int {
	if x == 0 || x == 1 {
		return big.NewInt(1)
	}
	result := big.NewInt(x)
	return result.Mul(result, fatorial(x-1))
}
func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}

// recursividade
func exemplo11() {
	fmt.Println(fatorial(100))

	//overflow
	fmt.Println(loops(100))
}
