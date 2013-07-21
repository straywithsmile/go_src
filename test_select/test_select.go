package main

import (
	"fmt"
)

func sub(c chan int, quit chan int) {
	c <- 3 + 2
	c <- 3 + 3
	c <- 0
	quit <- 0
}

func main() {
	channel := make (chan int)
	quit := make(chan int)
	go sub(channel, quit)
	for {
		select {
			case data := <- channel:
				fmt.Println(data)
				if data == 0 {
					fmt.Println("quit prev")
					//return
				}
			case <-quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("-----")
		}
	}
}
