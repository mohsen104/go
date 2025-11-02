package main

import (
	"fmt"
	"time"
)

func main() {
	numChan := make(chan int)

	go SendDataToChan(numChan)

	recivedNum := <-numChan
	PrintlnWithTime("Received Number:", recivedNum)

	time.Sleep(time.Second * 2)
}

func SendDataToChan(numChan chan int) {
	PrintlnWithTime("Before Sending 1 To Channel")
	numChan <- 1
	PrintlnWithTime("After Sending 1 To Channel")
	PrintlnWithTime("Before Sending 2 To Channel")
	numChan <- 2
	PrintlnWithTime("After Sending 2 To Channel")
	PrintlnWithTime("Before Sending 3 To Channel")
	numChan <- 3
	PrintlnWithTime("After Sending 3 To Channel")
}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
