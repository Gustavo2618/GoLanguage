package main

import "fmt"

//exercicios do cap 8
func main() {
	exemplo01()
	exemplo02()

}

//video 69
//exemplo de array
//exemplo slice e print de tipo
func exemplo01() {
	array := [10]int{1, 10, 100, 1000, 30, 40, 50, 60, 80, 99}

	for i := range array {
		fmt.Println(i, array[i])
	}
	fmt.Printf("%T", array)
}
//video 70
//exemplo slice e print de tipo
func exemplo02(){
	slice := []int{12, 101, 1008, 10000, 301, 402, 503, 604, 805, 996}

	for i := range slice {
		fmt.Println(i, slice[i])
	}
	fmt.Printf("%T", slice)
}