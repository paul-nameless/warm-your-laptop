package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	fmt.Println("Warming up your laptop...")
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {}
		}()
	}
	time.Sleep(2 * time.Minute)
	fmt.Println("Now your laptop is warm and cozy.")
}
