package main

import (
	"fmt"
	"math/rand"
)

// function to draw even or odd and compare with user pick
func evenOdd(bankroll, betAmount float64) {
	var userNum int
	var drawNum int
	allNum := []int{00}
	for i := 0; i < 37; i++ {
		allNum = append(allNum, i)
	}
	for userNum < 1 || userNum > 36 {
		fmt.Println("Please Pick a number (1-36): ")
		fmt.Scanln(&userNum)
	}
	drawNum = rand.Intn(len(allNum))
	if drawNum == 0 || drawNum == 1 {
		bankroll = bankroll - betAmount
		fmt.Println("Sorry Try again")
	} else if drawNum%2 == 0 && userNum%2 == 0 {
		bankroll = bankroll + betAmount
		fmt.Println("congragulations you won!")
	} else if drawNum%2 != 0 && userNum%2 != 0 {
		bankroll = bankroll + betAmount
		fmt.Println("congragulations you won!")
	} else {
		bankroll = bankroll - betAmount
		fmt.Println("Sorry Try again")
	}
}

// function to draw red or black and compare with user pick
func redBlack(bankroll, betAmount float64) {
	var userColor string
	fmt.Println("Please Pick your color(red or black): ")
	fmt.Scanln(&userColor)
	var colorInt int
	var drawColor string
	colorInt = rand.Intn(2)
	if colorInt == 0 {
		drawColor = "red"
	} else {
		drawColor = "black"
	}
	//fmt.Println(drawColor)
	if userColor == drawColor {
		bankroll = bankroll + betAmount
		fmt.Println("congragulations you won!")
	} else {
		bankroll = bankroll - betAmount
		fmt.Println("Sorry Try again")
	}
}

// function to draw a number and compare with user pick
func draw(bankroll, betAmount float64) {
	var userNum int
	var drawNum int
	allNum := []int{}
	for i := 0; i < 37; i++ {
		allNum = append(allNum, i)
	}
	drawNum = rand.Intn(38)
	for userNum < 1 || userNum > 36 {
		fmt.Println("Please Pick a number (0-36): ")
		fmt.Scanln(&userNum)
	}
	fmt.Println(userNum, drawNum)
	if drawNum == userNum {
		bankroll = bankroll + (betAmount * 35)
		fmt.Println("congragulations you won!")
	} else {
		bankroll = bankroll - betAmount
		fmt.Println("Sorry Try again")
	}
}

// func to pick high or low
func highLow(bankroll, betAmount float64) {
	var userChoice string
	var drawChoicerand int
	var drawChoice string
	fmt.Println("Pick your choice Low(0-17) or High(18-36)")
	fmt.Scanln(&userChoice)
	drawChoicerand = rand.Intn(2)
	if drawChoicerand == 0 {
		drawChoice = "Low"
	} else {
		drawChoice = "High"
	}
	if userChoice == drawChoice {
		bankroll = bankroll + betAmount
		fmt.Println("congragulations you won!")
	} else {
		bankroll = bankroll - betAmount
		fmt.Println("Sorry Try again")
	}

}

func checkBetAmount(bankroll, betAmount float64) {
	for betAmount < 1 {
		fmt.Println("Minimum bet amount is $1. Try again")
		fmt.Scanln(&betAmount)
	}
	for bankroll < betAmount {
		fmt.Println("Your Bet Amount(", betAmount, ") is more than bankroll(", bankroll, "). Pleaes check and try again")
		fmt.Scanln(&betAmount)
	}
}

func checkBetType() int {
	var betTypeNum int = -1
	for betTypeNum < 0 || betTypeNum > 5 {
		fmt.Println("Please choose your bet Type from below options")
		fmt.Println("1. red/black\n2. odd/even\n3. pick Num\n4. Low/High")
		fmt.Scanln(&betTypeNum)
	}
	return betTypeNum
}

// func to choose bet type
func betType(bankroll, betAmount float64) {
	checkBetAmount(bankroll, betAmount)
	var betTypeNum int = checkBetType()
	bankroll = bankroll - betAmount
	fmt.Println(bankroll)
	switch betTypeNum {
	case 1:
		redBlack(bankroll, betAmount)
	case 2:
		evenOdd(bankroll, betAmount)
	case 3:
		draw(bankroll, betAmount)
	case 4:
		highLow(bankroll, betAmount)
	default:
		fmt.Println("please choose correct bet type number from above")
	}
}

func main() {

	var bankroll float64  //user total bet amount
	var betAmount float64 //user bet amount for one bet

	fmt.Println("Please enter your Bankroll: ")
	fmt.Scanln(&bankroll)
	for bankroll < 1 {
		fmt.Println("Amount should be greater than $1. Please try again with appropriate amount")
		fmt.Scanln(&bankroll)
	}
	fmt.Println("Please enter your bet amount for the current roll")
	fmt.Scanln(&betAmount)
	betType(bankroll, betAmount)
}
