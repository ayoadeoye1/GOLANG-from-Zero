package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	rounds := 3
	fmt.Printf("The Game Begins, You have %d rounds to play \n", rounds)

	for i:=0; i < rounds; i++ {
		var computerNum = rand.Intn(rounds)
		var computerChoice string

		switch computerNum {
		case 0:
			computerChoice = "rock"
		case 1:
			computerChoice = "papper"
		case 2:
			computerChoice = "scissors"
		}

		var yourChoice string

		fmt.Print("Enter your choice and game (Rock Papper Scissors):")
		fmt.Scanln(&yourChoice)
		yourChoice = strings.ToLower(yourChoice)

		if yourChoice == computerChoice {
			fmt.Println("Game Tie")
		} else if
			yourChoice == "rock" && computerChoice == "scissors" ||
			yourChoice == "papper" && computerChoice == "rock" ||
			yourChoice == "scissors" && computerChoice == "papper" {
			fmt.Println("You Won This Round!")
		} else {
			fmt.Println("Computer Won!")
		}

		fmt.Printf("Round %d played! \n", i+1)
	}
	fmt.Println("Game Over!")
}