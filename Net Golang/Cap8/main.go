package main

import "fmt"

func main() {
	exemplo01()
	exemplo02()
	exemplo03()
}

// video59 array
func exemplo01() {
	var teste1 [5]int
	var teste2 [100]int

	teste1[0] = 1
	teste1[1] = 10
	fmt.Println(teste1[0], teste1[1])
	fmt.Println(teste1)
	fmt.Printf("%T\n", teste1)
	fmt.Printf("%T\n", teste2)
	//tamanho dos arrays
	fmt.Println(len(teste2))
}

//video60 slices
func exemplo02() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println("array: ", array)
	slice := []int{6, 7, 8, 9}
	fmt.Println("slice: ", slice)
	// array2 := append(array, 6)
	slice2 := append(slice, array[0])
	// fmt.Println("array2: ", array2)
	fmt.Println("slice2: ", slice2)
}

//video61
//função range percorre todo o slice
func exemplo03() {
	slice := []string{"banana", "maça", "jaca", "Pera", "morango", "Laranja", "Pêssego"}
	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor:", valor)
	}
	fmt.Println()
	//possivel erro de tamanho por conta do tamanho do slice
	// slice[7] = "melancia"
	//ignorando valores com _ ou indice ou valor
	slice = append(slice, "melancia")
	for indice, _ := range slice {
		fmt.Println("No indice", indice, "elementos")
	}
	fmt.Println()
	slice = append(slice, "melancia")
	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %s. \n", valor)
	}
	fmt.Println()
}
