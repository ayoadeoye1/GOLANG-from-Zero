package main

import "fmt"

func logMatrix(x int) {
	mainArr := make([][]int, 0)
	for i := 0; i < x; i++ {
		subArr := make([]int, 0)
		for j := 1; j < x; j++ {
			subArr = append(subArr, i*j)
		}
		mainArr = append(mainArr, subArr)
	}

	fmt.Print(mainArr)
}

func main() {
	logMatrix(10)
}
