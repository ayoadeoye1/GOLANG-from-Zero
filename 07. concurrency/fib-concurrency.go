package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func fibMain() {
	start := time.Now()
	bm := 15
	fibCh := make(chan float64, bm)

	for i := 1; i < bm; i++ {
		go fib(float64(i), fibCh)
		fmt.Printf("Fib(%v): %v\n", i, <-fibCh)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
