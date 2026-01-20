package main

import (
	"fmt"
	"time"
)

func somefunc(num string) {
	fmt.Println(num)
}

func main() { //in go main function is synchronous, if i want to make it asynchronous i have to use go keyword
	// in front of somefunc. in this case somefuc fork in it's own process and this is called goroutine
	// we can run it concurrently by go routine
	go somefunc("1")
	go somefunc("2")
	go somefunc("3")

	time.Sleep(time.Second * 2) //by waiting for 2 seconds somefunc will complete its execution before main function run

	fmt.Println("Main Function")
}
