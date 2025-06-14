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

//video62
//usando slices, pegando fatias e comando len
//exercicio
func exemplo04() {
	//sabores[:], sabores[a:],sabores[:b] sabores[a:b]
	sabores := []string{"pepperoni", "mussarela", "portuguesa", "calabresa", "marguerita", "quatro queijos"}
	fatia := sabores[0:2]
	fmt.Println(fatia)
	fatia1 := sabores[3:len(sabores)]
	fmt.Println(fatia1)
	fatia2 := sabores[1:3]
	fmt.Println(fatia2)

	//usando len
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
	//Println
	fmt.Println(sabores[:])

	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

//video63
//adicionando slices e os enumeration ...
func exemplo05() {
	vec1 := []int{1, 2, 3, 4, 5, 6}
	vec2 := []int{7, 8, 9, 10, 11, 12}
	fmt.Println(vec1)
	vec3 := append(vec1, 0, 2, 7, 9, 10)
	fmt.Println(vec3)
	vec3 = append(vec3, vec2...)
	fmt.Println(vec3)

}

//video 64
//usando slice com make
func exemplo06() {
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4
	slice[4] = 10
	//causa erro por conta do tamanho do slice
	//slice[5] = 50
	//para aumentar o slice usar append
	slice = append(slice, 50)
	slice = append(slice, 51)
	slice = append(slice, 52)
	slice = append(slice, 53)
	slice = append(slice, 54)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
	//aumentando capacidade do slice
	slice = append(slice, 55)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
	slice = append(slice, 56)
	fmt.Println("slice: ", slice, len(slice), cap(slice))
}

//video 65
//slices multidimencionais matrizes
func exemplo07() {
	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println("slice de slice: ", ss)
	fmt.Println("linha slice: ", ss[1])
	fmt.Println("item de ss: ", ss[2][1])

	for indice := range ss {
		for indece2 := range ss {
			fmt.Println("item de ss: ", ss[indice][indece2])
		}
	}

	ss1 := [][][][]int{
		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{20, 10, 30, 40, 50},
			},
		},
		[][][]int{
			[][]int{
				[]int{2, 4, 6, 8, 10},
			},
			[][]int{
				[]int{3},
			},
		},
	}
	fmt.Println(ss1[1])
	fmt.Println(ss1[1][0])
	fmt.Println(ss1[1][0][0])
	fmt.Println(ss1[1][0][0][2])

	for indices1 := range ss1 {
		for indeces2 := range ss1[indices1] {
			for indices3 := range ss1[indices1][indeces2] {
				for indeces4 := range ss1[indices1][indeces2][indices3] {
					fmt.Println("item de ss1: ", ss1[indices1][indeces2][indices3][indeces4])
				}
			}
		}
	}
}
