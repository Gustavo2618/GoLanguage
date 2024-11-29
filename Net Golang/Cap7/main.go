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

func exemplo01() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exemplo02() {
	for i := 65; i < 91; i++ {
		for k := 0; k < 3; k++ {
			fmt.Printf("\t%#U \n", i)
		}
	}
}

func exemplo03() {
	minhaIdade := 1991

	for minhaIdade < 2025 {
		fmt.Println(minhaIdade)
		minhaIdade++
	}
	fmt.Println()
}

func exemplo04() {
	minhaIdade := 1991

	for {
		if minhaIdade < 2024 {
			fmt.Println(minhaIdade)
			minhaIdade++
		} else {
			break
		}

	}
}

func exemplo05() {
	for i := 10; i < 100; i++ {
		rest := i % 4
		fmt.Printf("\nDivisão de %d por 4 tem resto: %d\n ", i, rest)
	}
}

func exemplo06() {
	for i := 10; i < 100; i++ {
		rest := i % 4
		if rest == 0 {
			fmt.Printf("\nDivisão de %d por 4 não tem resto: %d\n ", i, rest)
		} else if rest == 1 {
			fmt.Printf("\nDivisão de %d por 4 tem resto: %d\n ", i, rest)
		} else {
			fmt.Printf("\nDivisão de %d por 4 tem resto: %d diferente de 1 e 0\n ", i, rest)
		}
	}
}

func exemplo07() {

	esporteFavorito := "natação"
	switch {
	case esporteFavorito == "natação":
		fmt.Printf("\nMeu esporte é %s\n", esporteFavorito)
	case esporteFavorito == "futebol":
		fmt.Printf("\nMeu esporte é %s\n", esporteFavorito)
	case esporteFavorito == "basquete":
		fmt.Printf("\nMeu esporte é %s\n", esporteFavorito)
	default:
		fmt.Printf("\nVocê não pratica esporte....\n")
	}
}

func exemplo08() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || true)
	fmt.Println(!true)
}
