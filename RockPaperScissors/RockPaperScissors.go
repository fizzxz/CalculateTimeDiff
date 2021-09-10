package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	choices     = []string{"Rock", "Paper", "Scissors"}
	usersChoice int
	play        bool
)

func main() {
	play = true
	for play {

		fmt.Println("Enter one of the following: ")
		fmt.Println("Rock (0), Paper(1) or Scissors(2) or Quit (Any other value): ")

		_, err := fmt.Scanln(&usersChoice)
		if err != nil {
			fmt.Println(err)
			play = false
		}

		play = isInputValid(usersChoice)
		if play {
			rockPaperScissors(usersChoice)
		} else {
			fmt.Println("Game ended, thanks for playing!")
		}
	}
}

func rockPaperScissors(input int) {
	cpuVal := cpuMove()
	switch {
	case input == cpuVal:
		fmt.Println(choices[input] + " matches " + choices[cpuVal])
		fmt.Println("It's a tie!")

	case (input == 0 && cpuVal == 2) || (input == 1 && cpuVal == 0) || (input == 2 && cpuVal == 0):
		fmt.Println(choices[input] + " beats " + choices[cpuVal])
		fmt.Println("You Win Ez👏 😎")

	default:
		fmt.Println(choices[cpuVal] + " beats " + choices[input])
		fmt.Println("Computer wins!")
		fmt.Println("Better luck next time 😒👌")
	}
}

func cpuMove() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 2
	return (rand.Intn(max-min+1) + min)
}

func isInputValid(userInput int) bool {
	if userInput == 0 || userInput == 1 || userInput == 2 {
		return true
	} else {
		return false
	}
}
