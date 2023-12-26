package main

import (
	"connect4/internal"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCheckIfGameFinished(t *testing.T) {
	type test struct {
		count int
		board internal.Board
		want  int
	}

	var tests []test

	// read data from testBoards.txt
	// increase count
	// read board
	// read want

	data, err := os.ReadFile("/home/nils/GoProjects/connect4/data/testBoards.txt")
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(data), "\n")
	board := internal.NewBoard(7, 6)

	// if 1,2,c oder d then board complete
	// else fill board
	var testBoard internal.Board
	for i := 0; i < len(lines); i++ {
		testBoard = *board
		switch lines[i] {
		case ".":
			continue
		case "X":
			testBoard.Columns[i].PlaceToken(internal.Player2Token)
		case "O":
			testBoard.Columns[i].PlaceToken(internal.Player1Token)

		}

	}
	testBoard.Print()
	for _, tc := range tests {
		got := tc.board.CheckIfGameFinished()
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v in count: %d", tc.want, got, tc.count)
		}
	}

}
