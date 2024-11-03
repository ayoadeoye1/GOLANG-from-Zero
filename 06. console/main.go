package main

import "fmt"

func logMatrix(x int) {

	ch := make(chan [][]int)

	go func() {
		mainArr := make([][]int, 0)
		for i := 0; i < x; i++ {
			subArr := make([]int, 0)
			for j := 1; j < x; j++ {
				subArr = append(subArr, i*j)
			}
			mainArr = append(mainArr, subArr)
		}

		ch <- mainArr
	}()

	chanVal := <-ch

	fmt.Print(chanVal)
}

func main() {
	logMatrix(10)
}
