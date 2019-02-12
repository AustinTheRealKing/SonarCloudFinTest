// connect4.go for CSI 380 Assignment 3
// The struct C4Board should implement the Board
// interface specified in Board.go
// Note: you will almost certainly need to add additional
// utility functions/methods to this file.

package main

import "strings"

// size of the board
const numCols uint = 7
const numRows uint = 6

// size of a winning segment in Connect 4
const segmentLength uint = 4

// The main struct that should implement the Board interface
// It maintains the position of a game
// You should not need to add any additional properties to this struct, but
// you may add additional methods
type C4Board struct {
	position [numCols][numRows]Piece  // the grid in Connect 4
	colCount [numCols]uint // how many pieces are in a given column (or how many are "non-empty")
	turn     Piece // who's turn it is to play
}

// Who's turn is it?
func (board C4Board) Turn() Piece {
	return board.turn
}

// Put a piece in column col.
// Returns a copy of the board with the move made.
// Does not check if the column is full (assumes legal move).
func (board C4Board) MakeMove(col Move) Board {
	var temp C4Board = board
	temp.colCount[col] += 1
	return temp
}

// All of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {
	tempMove := []Move{}
	for i := 0; i < len(board.colCount); i++{
		if board.colCount[i] < numRows{
			tempMove= append(tempMove, Move(i))
		}
	}
	return tempMove
}


// Is it a win?
func (board C4Board) IsWin() bool {
	for i := 0; i < 6; i++{
		for j := 0; j < 7; j++{
				if i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 1)]{
					if i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 2)]{
						if i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 3)]{
							return true
						}
					}
				}
				if j + 1 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 1)][uint(i)]{
					if j + 2 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 2)][uint(i)]{
						if j + 3 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 3)][uint(i)]{
							return true
						}
					}
				}
				if j + 1 < 7 && i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 1)][uint(i + 1)]{
					if j + 2 < 7 && i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 2)][uint(i + 2)]{
						if j + 3 < 7 && i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 3)][uint(i + 3)]{
							return true
						}
					}
				}

				if j - 1 >= 0 && i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 1)][uint(i + 1)]{
					if j - 2 >= 0 && i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 2)][uint(i + 2)]{
						if j - 3 >= 0 && i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 3)][uint(i + 3)]{
							return true
						}
					}
				}
			}
	}
	return false
}

// Is it a draw?
func (board C4Board) IsDraw() bool {
	for i := 0;i < len(board.colCount) ; i++{
		if board.colCount[i] != 6 {
			return false
		}
	}
	if board.IsWin() {
		return false
	}else{
		return true
	}
}

// Who is winning in this position?
// This function scores the position for player
// and returns a numerical score
// When player is doing well, the score should be higher
// When player is doing worse, player's returned score should be lower
// Scores mean nothing except in relation to one another; so you can
// use any scale that makes sense to you
// The more accurately Evaluate() scores a position, the better that minimax will work
// There may be more than one way to evaluate a position but an obvious route
// is to count how many 1 filled, 2 filled, and 3 filled segments of the board
// that the player has (that don't include any of the opponents pieces) and give
// a higher score for 3 filleds than 2 filleds, 1 filleds, etc.
// You may also need to score wins (4 filleds) as very high scores and losses (4 filleds
// for the opponent) as very low scores
func (board C4Board) Evaluate(player Piece) float32 {
	var array [7][6]int
	score := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++{
			if board.position[uint(j)][uint(i)].String() == "+" && array[j][i] != 1 {
				if j + 1 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 1)][uint(i)] && array[j + 1][i] != 1 {
					array[j+1][i] = 1
					array[j][i] = 1
					score += 10
					if j + 2 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 2)][uint(i)] && array[j + 2][i] != 1 {
						array[j+2][i] = 1
						score += 5
						if j + 3 < 7 && board.position[uint(j)][uint(i)] == board.position[uint(j + 3)][uint(i)] && array[j + 3][i] != 1{
							array[j+3][i] = 1
							score += 1000
						}
					}
				}
			}

			if board.position[uint(j)][uint(i)].String() == "+" && array[j][i] != 2 {
				if i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 1)] && array[j][i + 1] != 2 {
					array[j][i] = 2
					array[j][i+1] = 2
					score += 10
					if i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 2)] && array[j][i + 2] != 2 {
						array[j][i + 2] = 2
						score += 5
						if i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j)][uint(i + 3)] && array[j][i + 3] != 2{
							array[j][i + 3] = 2
							score += 1000
						}
					}
				}
			}

			if board.position[uint(j)][uint(i)].String() == "+" && array[j][i] != 3 {
				if j + 1 < 7 && i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 1)][uint(i + 1)] && array[j + 1][i + 1] != 3{
					array[j][i] = 3
					array[j + 1][i + 1] = 3
					score += 10
					if j + 2 < 7 && i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 2)][uint(i + 2)] && array[j + 2][i + 2] != 3{
						array[j + 2][i + 2] = 3
						score += 5
						if j + 3 < 7 && i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j + 3)][uint(i + 3)] && array[j + 3][i + 3] != 3{
							array[j + 3][i + 3] = 3
							score += 1000
						}
					}
				}
			}

			if board.position[uint(j)][uint(i)].String() == "+" && array[j][i] != 4{
				if j - 1 >= 0 && i + 1 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 1)][uint(i + 1)] && array[j - 1][i + 1] != 4{
					array[j - 1][i + 1] = 4
					array[j][i] = 4
					score += 10
					if j - 2 >= 0 && i + 2 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 2)][uint(i + 2)] && array[j - 2][i + 2] != 4{
						array[j - 2][i + 2] = 4
						score += 5
						if j - 3 >= 0 && i + 3 < 6 && board.position[uint(j)][uint(i)] == board.position[uint(j - 3)][uint(i + 3)] && array[j - 3][i + 3] != 4{
							array[j - 3][i + 3] = 4
							score += 1000
						}
					}
				}
			}
		}
	}
	return float32(score)
}
/*
function return score(takes in the length of the connection and is a returns the number
SWITCH

If length == 1
	return 5
if length == 2
	return 10
if length == 3
	return 25
if length == 4
	return 1000
default case is 0

 */

// Nice to print board representation
// This will be used in play.go to print out the state of the position
// to the user
func (board C4Board) String() string {
	finalString := ""
	for i := 5; i >= 0; i--{
		finalString += "| "
		for j := 0; j < 7; j++{
			finalString = finalString + board.position[uint(j)][uint(i)].String() +  " | "
		}
		finalString += "\n"
	}

	finalString += "\n\n\n"
	return finalString
}
