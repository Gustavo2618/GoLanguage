package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

}

func CountTime(hours float64, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Println(start)
	var Second = (hours * 3600)
	duration := (time.Duration(Second) * time.Second)
	fmt.Println("Tempo total da go routine: ", duration)
	//time
	for time.Since(start) < duration {
		time.Sleep(time.Duration(Second) * time.Second)
		elapsed := time.Since(start)
		fmt.Printf("Tempo decorrido : %s\n", elapsed)
	}
	fmt.Println("Tempo total decorrido: ", time.Since(start))
}

func GoRoutineTime(CurrentRandom, CurrentInterval float64) {
	fmt.Println("starting count time function...")
	wg.Add(1)
	go CountTime(CurrentRandom, &wg)
	wg.Wait()
	fmt.Println("Notifing machines to do atestation....")
	// reminder := Reminder(CurrentInterval, CurrentRandom)
	// wg.Add(1)
	// go CountTime(reminder, &wg)
	// wg.Wait()
	// fmt.Println("We are ready for do another attestation period.")
	// NextRandom := random.GenerateRandomFloat(0, CurrentInterval)
	// teste := TimeForNextAttestation(NextRandom, reminder)
	// fmt.Println("-> `Proximo random: ", teste)
}
