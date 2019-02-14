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
	for true {
		if contains(gameBoard.LegalMoves(), userInput) {
			return userInput
		} else {
			if counter > 0 {
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
	var tempMove Move

	for i := 0; true; i++ {

		if i%2 == 0 {
			tempMove = getPlayerMove()
			fmt.Println("MY TURN!", gameBoard.Turn())
			gameBoard = gameBoard.MakeMove(tempMove)
			if gameBoard.IsWin() {
				fmt.Println("You won!")
				break
			}
		} else {
			tempMove = FindBestMove(gameBoard, 3)
			fmt.Println("OPPONENT TURN!", gameBoard.Turn())
			gameBoard = gameBoard.MakeMove(tempMove)
			if gameBoard.IsWin() {
				fmt.Println("You lost!")
				break
			}
		}
		if gameBoard.IsDraw() {
			fmt.Println("Its a draw!")
			break
		}
	}
	fmt.Println(gameBoard.String())
}
