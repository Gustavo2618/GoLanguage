// package main

// import (
// 	"fmt"
// )

// type Conta struct {
// 	saldo float64
// }

// func (c *Conta) depositarDezReais() float64 {
// 	// c.saldo += 10
// 	return c.saldo + 10
// }

// func main() {
// 	contaTeste := Conta{saldo: 10}

// 	contaTeste.depositarDezReais()
// 	fmt.Println(contaTeste.depositarDezReais())

// 	fmt.Println(contaTeste)
// }

// package main

// import (
// 	"fmt"
// )

// func Somando(numeros ...int) int {
// 	resultadoDaSoma := 0
// 	for _, numero := range numeros {
// 		resultadoDaSoma += numero
// 	}
// 	return resultadoDaSoma
// }

//	func main() {
//		fmt.Println(Somando(1))
//		fmt.Println(Somando(1, 1))
//		fmt.Println(Somando(1, 1, 1))
//		fmt.Println(Somando(1, 1, 2, 4))
//	}
// package main

// import (
// 	"fmt"
// )

// func SemParametro() string {
// 	return "Exemplo de função sem parâmetro!"
// }

// func UmParametro(texto string) string {
// 	return texto
// }

// func DoisParametros(texto string, numero int) (string, int) {
// 	return texto, numero
// }

// func main() {
// 	fmt.Println(SemParametro())
// 	fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
// 	fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))
// }

package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/go-co-op/gocron/v2"
)

// func main() {
// 	s, _ := gocron.NewScheduler()
// 	defer func() { _ = s.Shutdown }()

// 	_, _ = s.NewJob(
// 		gocron.DailyJob(
// 			1,
// 			gocron.NewAtTimes(
// 				gocron.NewAtTime(19, 25, 0),
// 				gocron.NewAtTime(14, 0, 0),
// 			),
// 		),
// 		gocron.NewTask(
// 			func(a, b string) {
// 				fmt.Println("escrevendo A e B: ", a, b)
// 			},
// 		),
// 	)
// 	s.Start()

// 	select {
// 	case <-time.After(time.Minute):
// 	}

// }
func main() {
	s, _ := gocron.NewScheduler()
	defer func() { s.Shutdown() }()

	j, _ := s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(time.Second*10)),
		),

		gocron.NewTask(
			func() {
				fmt.Println("Notify Machines to do Attestation...")
			},
		),
		gocron.WithTags("endpoint1"),
	)
	// fmt.Println("testando job: ", j.ID())
	fmt.Println("Cheguei aqui no start!")

	// teste = append(j.Tags(), id)
	// fmt.Println(j.ID())
	s.Start()
	fmt.Println(j.Tags())
	select {
	case <-time.After(time.Second * 11):
	}
	fmt.Println("Cheguei aqui no fim do start!")
	fmt.Println(j.Tags())

}
func CountTime(hours float64, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	// fmt.Println(start)
	var Second = (hours * 3600)
	duration := (time.Duration(Second) * time.Second)
	// fmt.Println(duration)
	//time
	for time.Since(start) < duration {
		time.Sleep(time.Duration(Second) * time.Second)
		// elapsed := time.Since(start)
		// fmt.Printf("Tempo decorrido : %s\n", elapsed)
	}
	// fmt.Println("Tempo total decorrido: ", time.Since(start))
}

//  go mod init codigo.go //iniciando mod file
// go get -a  //importando pacotes do github

//create a scheduler
// s, err := gocron.NewScheduler()
// if err != nil {
// 	fmt.Println("deu erro!")
// }

// j, err := s.NewJob(
// 	gocron.DurationJob(time.Minute),
// 	// gocron.NewTask(
// 	// 	func(a string, b int) {

// 	// 		fmt.Println("imprimindo:", a, b)
// 	// 	}, "hello", 1),

// 	gocron.NewTask(func(hours float64, wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		start := time.Now()
// 		Seconds := (hours * 3600)
// 		duration := (time.Duration(Seconds) * time.Second)
// 		fmt.Println(duration)
// 		for time.Since(start) < duration {
// 			time.Sleep(time.Duration(Seconds) * time.Second)
// 		}
// 		fmt.Println("Tempo total decorrido: ", time.Since(start))
// 	}, 0.003, &wg),
// )

// if err != nil {
// 	fmt.Println("deu erro!")
// }

// fmt.Println("Escrevendo o id: ", j.ID())

// s.Start()

// select {
// case <-time.After(time.Minute):
// }

// err = s.Shutdown()

// if err != nil {
// 	fmt.Println("deu erro!")
// }
