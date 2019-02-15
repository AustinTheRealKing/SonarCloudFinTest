// minimax.go for CSI 380 Assignment 3
// This file contains a working implementation of Minimax
// You will need to implement the FindBestMove() methods to
// actually evaluate a position by running MiniMax on each of the legal
// moves in a starting position and finding the move associated with the best outcome
package main

import (
	"math"
	"sync"
)

type MoveAndEval struct {
	move Move
	eval float32
}

// Find the best possible outcome evaluation for originalPlayer
// depth is initially the maximum depth
func MiniMax(board Board, maximizing bool, originalPlayer Piece, depth uint) float32 {
	// Base case — terminal position or maximum depth reached
	if board.IsWin() || board.IsDraw() || depth == 0 {
		return board.Evaluate(originalPlayer)
	}

	// Recursive case - maximize your gains or minimize the opponent's gains
	if maximizing {
		var bestEval float32 = -math.MaxFloat32 // arbitrarily low starting point
		for _, move := range board.LegalMoves() {
			result := MiniMax(board.MakeMove(move), false, originalPlayer, depth-1)
			if result > bestEval {
				bestEval = result
			}
		}
		return bestEval
	} else { // minimizing
		var worstEval float32 = math.MaxFloat32
		for _, move := range board.LegalMoves() {
			result := MiniMax(board.MakeMove(move), true, originalPlayer, depth-1)
			if result < worstEval {
				worstEval = result
			}
		}
		return worstEval
	}
}

// Find the best possible move in the current position
// looking up to depth ahead
// This version looks at each legal move from the starting position
// concurrently (runs minimax on each legal move concurrently)
func ConcurrentFindBestMove(board Board, depth uint) Move {
	//https://stackoverflow.com/questions/18499352/golang-concurrency-how-to-append-to-the-same-slice-from-different-goroutines?fbclid=IwAR0BmXxVgW6vWVtdVf42fswFzuznJzt4nvuhQNGHnI32ZA-uWRyO8sW0H0A
	//https://stackoverflow.com/questions/46010836/using-goroutines-to-process-values-and-gather-results-into-a-slice?fbclid=IwAR2aLUbrywI7Y6UWluYof0vZWHytDoMX5CWQpU5Sd6aITJWl5k0IM0V9TBM
	//Help with Concurrency from Sources above

	var wg sync.WaitGroup
	var sliceOfMoves []MoveAndEval
	var bestMove float32 = -math.MaxFloat32

	channel := make(chan MoveAndEval)
	for _, move := range board.LegalMoves() {
		wg.Add(1)
		go func(move Move) {
			defer wg.Done()
			miniMaxEval := MiniMax(board.MakeMove(move), false, board.Turn(), depth)
			channel <- MoveAndEval{move, miniMaxEval}
		}(move)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for s := range channel {
		sliceOfMoves = append(sliceOfMoves, s)
	}
	index := 0
	for i := 0; i < len(sliceOfMoves); i++ {
		if sliceOfMoves[i].eval > bestMove {
			index = i
			bestMove = sliceOfMoves[i].eval
		}
	}
	return sliceOfMoves[index].move

	/*
		var allPossibleMoves = board.LegalMoves()
		var bestMoveEval float32 = -math.MaxFloat32
		var bestMove Move

		channel := make(chan MoveAndEval)

		concurrentRun := func(board Board, move Move, channel chan MoveAndEval) {
			miniMaxEval := MiniMax(board, false, board.Turn(), depth)
			conncurentEval := MoveAndEval{move: move, eval: miniMaxEval}
			channel <- conncurentEval
		}

		for i := 0; i < len(allPossibleMoves); i++ {
			go concurrentRun(board.MakeMove(allPossibleMoves[i]), allPossibleMoves[i], channel)
		}

		for i := 0; i < len(allPossibleMoves); i++ {
			moveReturned := <-channel
			if moveReturned.eval > bestMoveEval {
				bestMoveEval = moveReturned.eval
				bestMove = moveReturned.move

			}
		}
		fmt.Println(bestMove)
		return bestMove
	*/

}

// Find the best possible move in the current position
// looking up to depth ahead
// This is a non-concurrent version that you may want to test first
func FindBestMove(board Board, depth uint) Move {
	var allPossibleMoves = Board.LegalMoves(board)
	var indexOfBestMove = 0
	var bestMove float32 = -math.MaxFloat32

	for i := 0; i < len(allPossibleMoves); i++ {
		miniMaxEval := MiniMax(board.MakeMove(allPossibleMoves[i]), false, board.Turn(), depth)
		if miniMaxEval > bestMove {
			indexOfBestMove = i
			bestMove = miniMaxEval
		}
	}

	return allPossibleMoves[indexOfBestMove]

}
