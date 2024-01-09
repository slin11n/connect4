package main

import (
	"connect4/internal"
	tm "github.com/buger/goterm"
)

func main() {
	board := internal.NewBoard(7, 6)
	board.Columns[1].PlaceToken(internal.Player2Token)
	board.Columns[1].PlaceToken(internal.Player2Token)
	board.Columns[1].PlaceToken(internal.Player2Token)
	board.Columns[1].PlaceToken(internal.Player2Token)
	board.Columns[6].PlaceToken(internal.Player1Token)
	//	board.Columns[3].PlaceToken(internal.Player1Token)
	//	board.Columns[5].PlaceToken(internal.Player1Token)
	board.Print()
	tm.MoveCursor(0, 9)
	tm.Print()
	if board.IsBoardFull() == true {
		tm.Println("Niemand hat gewonnen")
	}

	if board.CheckIfGameIsFinished(0, 0, "O", 0, "h") == internal.Player1 {
		tm.Println("O hat gewonnen")
	}
	if board.CheckIfGameIsFinished(0, 0, "X", 0, "h") == internal.Player2 {
		tm.Println("X hat gewonnen")
	}

	if board.CheckIfGameIsFinished(0, 0, "O", 0, "v") == internal.Player1 {
		tm.Println("O hat gewonnen")
	}
	if board.CheckIfGameIsFinished(0, 0, "X", 0, "v") == internal.Player2 {
		tm.Println("X hat gewonnen")
	}

	if board.CheckIfGameIsFinished(0, 0, "O", 0, "d") == internal.Player1 {
		tm.Println("O hat gewonnen")
	}
	if board.CheckIfGameIsFinished(0, 0, "X", 0, "d") == internal.Player2 {
		tm.Println("X hat gewonnen")
	}

	if board.IsBoardFull() == true {
		tm.Println("Niemand hat gewonnen")
	}

}
