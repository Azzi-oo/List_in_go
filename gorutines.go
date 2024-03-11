package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go myFunc()
	time.Sleep(1 * time.Second)

	runtime.GOMAXPROCS(4)
	go func() {
		fmt.Println("Привет, я анонимная горутина")
	}()
	time.Sleep(1 * time.Second)
}

func myFunc() {
	fmt.Println("hello")
}
