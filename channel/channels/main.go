package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}


func main() {

	// Channel E.g.
	// dataChan := make(chan int, 1) // Space for 1 answer
	dataChan := make(chan int)

	// go func() {
	// 	dataChan <- 996
	// }()
	//dataChan <- 996
	//n := <- dataChan // Push to var
	//fmt.Printf("n = %d \n", n)
	
	go func() {
		wg := sync.WaitGroup{}
		for i := 1; i < 1000000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n = %d \n", n)
	}


}