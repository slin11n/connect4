package main

import (
	"connect4/internal"
	"fmt"
	tm "github.com/buger/goterm"
)

func main() {
	board := internal.NewBoard(7, 6)
	board.Columns[0].PlaceToken(internal.Player1Token)
	board.Columns[2].PlaceToken(internal.Player1Token)
	board.Columns[3].PlaceToken(internal.Player1Token)
	board.Columns[4].PlaceToken(internal.Player1Token)
	board.Columns[5].PlaceToken(internal.Player1Token)
	//	board.Columns[3].PlaceToken(internal.Player1Token)
	//	board.Columns[5].PlaceToken(internal.Player1Token)
	result := board.CheckIfGameFinishedHorizontal()
	board.Print()
	tm.MoveCursor(0, 9)
	tm.Print()
	if result == internal.Player2 {
		fmt.Println("Spieler 2 hat gewonnen")
	}
	if result == internal.Player1 {
		fmt.Println("Spieler 1 hat gewonnen")
	}

}
