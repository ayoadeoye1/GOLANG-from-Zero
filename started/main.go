package main

import (
	"fmt"
	"time"
)
var count int32
func loop(n int){
	for i :=4; i<=n; i++ {
		if i%2!=0 && i%3!=0 && i%5!=0 && i%7!=0 {
			fmt.Println(i)
			count ++
		}
	}
}

func swic(rate uint){
	switch rate{
	case 1:
		fmt.Println("worse rating")
	case 2:
		fmt.Println("bad rating")
	case 3:
		fmt.Println("fair rating")
	case 4:
		fmt.Println("good rating")
	case 5:
		fmt.Println("excellent")
	default:
		fmt.Println("out of band")
	}
}

func main(){
	n := 100
	go loop(n)
	go swic(4)
	time.Sleep(time.Second)

	fmt.Println("hello word!");
	fmt.Printf("intergers that are divisible only by 1 and itself between 0 and %d is: %d \n ...", n, count);
}

