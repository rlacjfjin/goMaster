package learnConcurrency

import (
	"flag"
	"fmt"
	"time"
)

func function(){
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func UsingGoroutine(){
	go function()				//创建方式1

	go func() {						//创建方式2
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println()
}


func CreateMulti(){
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}