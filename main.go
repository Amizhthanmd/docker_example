package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for {
		go func() {
			for {
				fmt.Println("docker example")
				time.Sleep(1 * time.Second)
				break
			}
		}()
		time.Sleep(1 * time.Second)
		fmt.Println("no of go routines --->", runtime.NumGoroutine())
	}
}
