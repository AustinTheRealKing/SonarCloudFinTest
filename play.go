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
	/*
	HAHAHAHAHA I LOVE NOT HAVING GO WORK ON MY DESKTOP AND I HAVE TO BE ON DUTY

	Board gameBoard = new Board();
	Move tempMove;
	 
	for i = 0; true; i++
	{
		fmt.Print(gameBoard.String())
		
		if(i%2 == 0) //PlayersMove
		{
			tempMove = getPlayerMove()
			gameBoard = gameBoard.MakeMove(tempMove)
			if(gameBoard.isWin())
			{
				fmt.Print("You won!")
				break;
			}
		} else {

			tempMove = FindBestMove(gameBoard, 3)
			gameBoard = gameBoard.MakeMove(tempMove)
			if(gameBoard.isWin())
			{
				fmt.Print("You lost!")
				break;
			}
		}

		if(gameBoard.isDraw())
		{
			fmt.Print("Its a draw!")
			break;
		} 
	}
}
	*/
}
