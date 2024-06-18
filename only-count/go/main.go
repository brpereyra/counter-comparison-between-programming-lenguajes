package main

import (
	"fmt"
	"time"
)

func main() {
	// time comparison
	t1 := time.Now()
	var counter int
	for i := 0; i < 1000000; i++ {
		counter++
	}
	t2 := time.Since(t1)
	fmt.Println("Time take in microseconds: ", t2.Microseconds())
}
