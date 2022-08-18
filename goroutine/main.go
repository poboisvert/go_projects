package main

import "fmt"

func main() {
	oneChannel := make(chan int)
	twoChannel := make(chan int)

	go func() {
		oneChannel <- 1
	}()

	go func() {
		twoChannel <- 666
	}()

	msg := <- oneChannel
	fmt.Println(msg)

	select {
		case msgFromOneChannel := <- oneChannel:
			fmt.Println(msgFromOneChannel)
		
		case msgFromTwoChannel := <- twoChannel:
			fmt.Println(msgFromTwoChannel)
		}
}