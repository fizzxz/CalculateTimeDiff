package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	number      = cpuRandNum()
	prevNum     = 0
	score       = 0
	play        = true
	usersChoice string
)

func main() {
	for play {
		fmt.Println("The last number was: " + fmt.Sprint(number))
		fmt.Println("Enter plus (+) if you think the next number will be Higher")
		fmt.Println("or minus (-) if you think it will be Lower: ")
		_, err := fmt.Scanln(&usersChoice)
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to register a correct input")
			fmt.Println("Exiting app, final score = " + fmt.Sprint(score))
			play = false
		}

		if usersChoice == "+" || usersChoice == "-" {
			score += higherOrLower()
			fmt.Println("Current score: " + fmt.Sprint(score))
		} else {
			fmt.Println("failed to register a correct input")
			fmt.Println("Exiting app, final score = " + fmt.Sprint(score))
			play = false
		}
	}
}

func higherOrLower() int {
	prevNum = number
	number = cpuRandNum()
	if number < prevNum {
		if usersChoice == "+" {
			fmt.Println("You lose")
			return 0
		} else if usersChoice == "-" {
			fmt.Println("You Win")
			return 1
		}
	} else if number > prevNum {
		if usersChoice == "+" {
			fmt.Println("You Win")
			return 1
		} else if usersChoice == "-" {
			fmt.Println("You lose")
			return 0
		}
	} else if number == prevNum {
		fmt.Println("The random number generated is equal to the previous number, try again!")

		return 0
	}
	return 0
}

func cpuRandNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return (rand.Intn(max-min+1) + min)
}
