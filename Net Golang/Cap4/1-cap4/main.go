package main

import "fmt"

func main() {
	//uso de ""e para strings sem formatos quebra de linha etc
	//o acento agudo `` em uma string formata o texto da forma literal exemplo abaixo
	s := "Hello, playground"
	d := `hello,
	        mofi
	 
	  go playground!`

	m := []byte(s)
	fmt.Printf("%v\n%T\n", s, s)
	fmt.Printf("%v\n%T\n", d, d)
	fmt.Printf("%v\n%T\n", m, m)
	//caractere por caractere
	// for _, v := range s {
	// 	fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	// }
	// fmt.Println("")

	// //byte por byte
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	// }

	name := "ascii, é@â,  u99999"

	for _, x := range name {
		fmt.Printf("%v - %T - %#U - %#x\n", x, x, x, x)
	}
	fmt.Println("")

	//byte por byte
	for k := 0; k < len(name); k++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[k], s[k], s[k], s[k])
	}
}
