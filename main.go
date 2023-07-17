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

func main() {

	var bankroll float64  //user total bet amount
	var betAmount float64 //user bet amount for one bet
	var betType int       //user bet type

	fmt.Println("Please enter your Bankroll: ")
	fmt.Scanln(&bankroll)
	fmt.Println("Please enter your bet amount for the current roll")
	fmt.Scanln(&betAmount)
	if bankroll < betAmount {
		fmt.Println("Your Bet Amount is more than bankroll. Pleaes check and try again")
		fmt.Println(bankroll)
	} else {
		fmt.Println("Please choose your bet Type from below options")
		fmt.Println("1. red/black\n2. odd/even\n3. pick Num")
		fmt.Scanln(&betType)
		bankroll = bankroll - betAmount
	}
	switch betType {
	case 1:
		redBlack(bankroll, betAmount)
	case 2:
		evenOdd(bankroll, betAmount)
	case 3:
		draw(bankroll, betAmount)
	}

}
