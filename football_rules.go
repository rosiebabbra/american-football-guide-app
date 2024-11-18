package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Menu options
func showMenu() {
	fmt.Println("\nWelcome to the American Football Rules Guide!")
	fmt.Println("Please choose an option:")
	fmt.Println("1. What is American football?")
	fmt.Println("2. Objectives of the game")
	fmt.Println("3. Basic rules")
	fmt.Println("4. Scoring system")
	fmt.Println("5. Exit")
}

// Information about American football
func whatIsFootball() {
	fmt.Println("\nAmerican football is a team sport played on a rectangular field with goalposts at each end. Two teams of 11 players compete to score points by advancing the ball into the opposing team's end zone.")
}

// Objectives of the game
func objectivesOfGame() {
	fmt.Println("\nThe objective of American football is to score more points than the opposing team by the end of the game. This is done by moving the ball down the field and into the opposing team's end zone, either by running or passing the ball.")
}

// Basic rules of the game
func basicRules() {
	fmt.Println("\nBasic Rules:")
	fmt.Println("1. Each team has four downs (attempts) to advance the ball 10 yards.")
	fmt.Println("2. If the team fails to advance 10 yards, possession goes to the opposing team.")
	fmt.Println("3. The game is divided into four quarters of 15 minutes each.")
	fmt.Println("4. Players can pass the ball forward only once during a play, and it must be from behind the line of scrimmage.")
}

// Scoring system
func scoringSystem() {
	fmt.Println("\nScoring System:")
	fmt.Println("1. Touchdown: 6 points - Scored by carrying the ball into the opponent's end zone or catching it there.")
	fmt.Println("2. Extra Point: 1 point - Kicked through the goalposts after a touchdown.")
	fmt.Println("3. Two-Point Conversion: 2 points - An alternative to an extra point by running or passing the ball into the end zone.")
	fmt.Println("4. Field Goal: 3 points - Kicked through the goalposts during play.")
	fmt.Println("5. Safety: 2 points - Awarded to the defense when the offense is tackled in their own end zone.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		showMenu()
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			whatIsFootball()
		case "2":
			objectivesOfGame()
		case "3":
			basicRules()
		case "4":
			scoringSystem()
		case "5":
			fmt.Println("Thanks for learning about American football. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

