package main

import "fmt"

//exercicios do cap 8
func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
	exemplo05()
	exemplo06()

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
func exemplo02() {
	slice := []int{12, 101, 1008, 10000, 301, 402, 503, 604, 805, 996}

	for i := range slice {
		fmt.Println(i, slice[i])
	}
	fmt.Printf("%T", slice)
}

//video 71
//exemplo Printando partes especificas dos slices
func exemplo03() {
	slice := []int{12, 101, 1008, 10000, 301, 402, 503, 604, 805, 996}
	fmt.Println()
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}

//video 72
//incrementando slice
func exemplo04() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice = append(slice, 52)
	slice = append(slice, 53, 54, 55)
	fmt.Println("slice x:", slice[:])
	slice1 := []int{56, 57, 58, 59, 60}
	slice = append(slice, slice1...)
	fmt.Println("slice x:", slice[:])
}

//video 73
// fatiando slices
func exemplo05() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice = append(slice[:3], slice[6:]...)
	fmt.Println(slice)
}

//video 74
//criando slices com make mostrando cap e len
func exemplo06() {
	slice := make([]string, 0, 24)
	slice = append(slice, "Acre", "Alagoas", "Amapá", "Amazonas")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Bahia", "Ceará", "Espirito Santo", "Goiás")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Pará", "Paraíba", "Paraná", "Pernambuco")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Rondônia", "Roraima", "Santa Catarina", "São Paulo")
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, "Sergipe", "Tocantis")
	fmt.Println(slice, len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

//video 75
func exemplo07() {

}

func exemplo08() {

}
func exemplo09() {

}
func exemplo10() {

}
