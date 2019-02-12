// main.go for CSI 380 Assignment 3
// This file includes the main game loop
// that actually creates a human vs computer game.

package main

import "fmt"

var gameBoard Board = C4Board{turn: Black}

// Find the user's next move
func getPlayerMove() Move {
	var userInput Move = 9
	counter := 0
	fmt.Println(gameBoard.LegalMoves())
	fmt.Println(gameBoard.String())
	for true{
		if contains(gameBoard.LegalMoves(), userInput){
			return userInput
		} else {
			if counter > 0{
				fmt.Println("sorry ", userInput, " was not a valid column number")
				fmt.Println("valid moves are:  ", gameBoard.LegalMoves())
			}
			fmt.Print("Enter The Column Number: ")
			fmt.Scanln(&userInput)
		}
		counter++
	}
	return userInput
}

// Main game loop
func main() {
	// YOUR CODE HERE
}
