package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter the Number of temprature: ")
	var number float64
	fmt.Scanln(&number)
	number = float64(number)

	fmt.Print("Enter the unit to convert from: ")
	var unitFrom string
	fmt.Scanln(&unitFrom)
	unitFrom = strings.ToUpper(unitFrom)

	fmt.Print("Enter the unit converting to: ")
	var unitTo string
	fmt.Scanln(&unitTo)
	unitTo = strings.ToUpper(unitTo)

	switch unitFrom {
	case "C":
		if unitTo == "K" {
			calc := number + 273.15
			fmt.Printf("The Result Is: %f \n", calc)
		} else if unitTo == "F" {
			calc := number*9/5 + 32
			fmt.Printf("The Result Is: %f \n", calc)
		}
	case "F":
		if unitTo == "K" {
			calc := float64(number - 32*5/9 + 273.15)
			fmt.Printf("The Result Is: %f \n", calc)
		} else if unitTo == "C" {
			calc := float64(number - 32*5/9 + 32)
			fmt.Printf("The Result Is: %f \n", calc)
		}
	case "K":
		if unitTo == "C" {
			calc := float64(number - 273.15)
			fmt.Printf("The Result Is: %f \n", calc)
		} else if unitTo == "F" {
			calc := float64(number - 273.15*9/5 + 32)
			fmt.Printf("The Result Is: %f \n", calc)
		}
	default:
		fmt.Print("Invalid Input \n")
	}
}
