package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	//contador := 0
	var contador int64
	totaldegoroutines := 15
	var wg sync.WaitGroup
	//var mu sync.Mutex

	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			//mu.Lock()
			atomic.AddInt64(&contador, 1)
			//v := contador
			runtime.Gosched()
			//v++
			//contador = v
			fmt.Println("Contador: \t", atomic.LoadInt64(&contador))
			//mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Valor Final:", contador)
	fmt.Println("Goroutines", runtime.NumGoroutine())
}
