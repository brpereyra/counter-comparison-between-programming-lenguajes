package main

import (
	"fmt"
	"time"
)

func main() {
	// time comparison
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
	}
	t2 := time.Since(t1)
	fmt.Println("Time take in seconds: ", t2.Seconds())
}
