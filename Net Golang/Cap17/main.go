package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// exercicios do cap 16
func main() {
	exemplo01()
	exemplo02()
	exemplo03()
	exemplo04()
}

type user struct {
	First string
	Age   int
}

// exercicio video 120
func exemplo01() {
	u1 := user{
		"James",
		32,
	}
	u2 := user{
		"Penny",
		27,
	}
	u3 := user{
		"Mofi",
		54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	//video 121
	sb2 := `[{"First":"James","Age":32},{"First":"Penny","Age":27},{"First":"Mofi","Age":54}]`
	var user1 []user
	err1 := json.Unmarshal([]byte(sb2), &user1)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Printf("%+v", user1)
}

// exercicio video 122
type user1 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func exemplo02() {
	u1 := user1{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user1{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user1{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user1{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println(err)
	}

}

// video 123
func exemplo03() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

// video 124
type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// implementação por Age
type OrdenarPorIdade []user2

func (x OrdenarPorIdade) Len() int {
	return len(x)
}
func (x OrdenarPorIdade) Less(i, j int) bool {
	return x[i].Age < x[j].Age
}

func (x OrdenarPorIdade) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type OrdenarPorLastName []user2

func (x OrdenarPorLastName) Len() int {
	return len(x)
}
func (x OrdenarPorLastName) Less(i, j int) bool {
	return x[i].Last < x[j].Last
}

func (x OrdenarPorLastName) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type OrdenarPorSayings []user2

func (x OrdenarPorSayings) Len() int {
	return len(x)
}
func (x OrdenarPorSayings) Less(i, j int) bool {
	return x[i].Last < x[j].Last
}

func (x OrdenarPorSayings) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func exemplo04() {
	u1 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user2{u1, u2, u3}

	fmt.Println("Inicial:\n", users)
	sort.Sort(OrdenarPorIdade(users))
	fmt.Println("Age:\n", users)
	fmt.Println("Inicial:\n", users)
	sort.Sort(OrdenarPorLastName(users))
	fmt.Println("LastName:\n", users)
	fmt.Println("Inicial:\n", users)
	sort.Sort(OrdenarPorSayings(users))
	fmt.Println("Sayings:\n", users)

}

func exemplo05() {

}
