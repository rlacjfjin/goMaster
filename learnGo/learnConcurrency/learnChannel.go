package learnConcurrency

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func UsingChannel(){
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	//
	willClose := make(chan int, 10)
	willClose <- -1
	willClose <- 0
	willClose <- 2
	<-willClose
	<-willClose
	<-willClose
	close(willClose)
	read := <-willClose
	fmt.Println(read)

}


