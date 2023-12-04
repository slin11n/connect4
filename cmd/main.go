package main

import "connect4/internal"

func main() {
	board := internal.NewBoard(7, 6)
	board.Slots[3].PlaceToken(internal.Player1Token)
	board.Slots[4].PlaceToken(internal.Player2Token)
	board.Slots[2].PlaceToken(internal.Player1Token)
	board.Slots[1].PlaceToken(internal.Player2Token)
	board.Slots[4].PlaceToken(internal.Player1Token)
	board.Slots[4].PlaceToken(internal.Player2Token)
	board.Slots[2].PlaceToken(internal.Player1Token)
	board.Print()
}
