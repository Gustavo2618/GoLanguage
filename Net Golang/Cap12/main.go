package main

import "fmt"

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
func exemplo03() {

}
func exemplo04() {

}
func exemplo05() {

}
func exemplo06() {

}
func exemplo07() {

}
