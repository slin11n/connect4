package main

import (
	"connect4/internal"
	"fmt"
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
	resultHor := board.CheckIfGameFinishedHorizontal()
	board.Print()
	tm.MoveCursor(0, 9)
	tm.Print()
	if resultHor == internal.Player2 {
		fmt.Println("Spieler 2 hat gewonnen")
	}
	if resultHor == internal.Player1 {
		fmt.Println("Spieler 1 hat gewonnen")
	}
	resultVer := board.CheckIfGameFinishedVertical()
	if resultVer == internal.Player2 {
		fmt.Println("Spieler 2 hat gewonnen")
	}
	if resultVer == internal.Player1 {
		fmt.Println("Spieler 1 hat gewonnen")
	}

}
