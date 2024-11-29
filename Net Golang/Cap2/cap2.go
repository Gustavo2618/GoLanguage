package main

import "fmt"

//var é usado fora de escopo pode declarar variaveis globais, variaveis dentro de escopo sao declaradas com marmota :=
var x = 10

var a int
var b float64
var c string
var d bool

type hotdog int

var hot hotdog = 101

func main() {
	//fmt retorna numero de bytes(tamanho de todo o print) e algum erro se acontecer....
	//lembrar que o _ nao utiliza as informação da variavel suprimida.
	numerodebytes, _ := fmt.Println("Hello World!" + "first things, first!")
	fmt.Println(numerodebytes)

	//Operador :=  é um operador de pelo menos uma variavel nova(declaração de variavel)
	// e funciona apenas nesse bloco de código, dentro de um escopo.
	//Operador =  é um operador de atribuição
	// x = 20 | "x recebe 20"
	x := 10
	y := "mofi mofi"
	//usando Printf para printar a variavel e seu tipo com %T
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

	//go aceita o uso de varios operandos na mesma linha como mostrado a seguir.
	//o var atribui qualquer tipo de valor a variavel apenas uma vez
	z := 10 < 20
	fmt.Println("\n", z)
	fmt.Println(">>>>>>>Variaveis inicializadas <<<<<<")
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)

	//testando prints
	testePrint := "Meu nome e gustavo!"
	testeRawPrint := `"Meu nome e Gustavo"`
	fmt.Println(testePrint)
	fmt.Print("Teste String normal:", testePrint, "Teste Raw String: ", testeRawPrint, "\n")
	fmt.Printf(testePrint)

	//testando Sprint
	//Sprint sendo usado para armazenar strings em um variavel.
	testeSprint := "oi"
	testeSprint1 := "Bom dia"
	sprint := fmt.Sprint(testeSprint, " ", testeSprint1)
	fmt.Println(sprint)

	//variaveis criadas com {type hotdog int} mesmo tendo os mesmo valores nao podem ser atribuidas aos tipos primitivos.
	fmt.Printf("%T ", hot)

	//Convertendo variáveis
	vars := 10
	fmt.Printf("%v\n", vars)
	fmt.Printf("%v\n", vars)

	vars = int(hot)
	fmt.Printf("%v\n", vars)

}
