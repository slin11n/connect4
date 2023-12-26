package internal

import (
	tm "github.com/buger/goterm"
)

const (
	Player1 = iota
	Player2
	Draw
	NoResult
)

type Board struct {
	Columns     []Column // List of columns for the board
	ColumnCount int      // How many columns are used
	NextPlayer  int      // Which player is up next
}

func NewBoard(columns, capacity int) *Board {
	var cs []Column
	for i := 0; i < columns; i++ {
		c := NewColumn(capacity)
		cs = append(cs, *c)
	}
	return &Board{
		Columns:     cs,
		ColumnCount: columns,
		NextPlayer:  Player1,
	}
}

func (b *Board) Print() {
	tm.Clear()
	for x, c := range b.Columns {
		if c.Used == 0 {
			continue
		}
		for i := 0; i < c.Used; i++ {
			tm.MoveCursor(x+1, c.Capacity-i)
			token := c.GetToken(i)
			tm.Print(token)
		}
	}
	tm.MoveCursor(0, 7) // todo fix later
	tm.Println("0123456")
	tm.Flush()

}

func (b *Board) debugPrint() {

}
func (b *Board) CheckIfGameFinished() int {
	//result := b.CheckIfGameFinishedHorizontal()
	//result = b.CheckIfGameFinishedVertical()
	//result = b.CheckIfGameFinishedDiagonal()
	return 0
}

func (b *Board) GetMiddleColumn() *Column {
	index := b.ColumnCount / 2
	return &b.Columns[index]
}

/* checkIfGameFinishedHorizontal checks if someone won horizontaly
func (b *Board) CheckIfGameFinishedHorizontal() int {
	// nimm die mittlere spalte
	// von unten nach oben, von 0 ...
	// wenn da ein token liegt
	//   dann prüfe nach links, wieviele vom gleichen token gibt es
	//   dann prüfe nach rechts, wieviele vom gleichen token gibt es
	//   sind das 4?
	column := b.GetMiddleColumn()
	for i := 0; i < column.Capacity; i++ {
		token := column.GetToken(i)
		var counter int
		if token == "X" || token == "O" {
			for j := 0; j < b.ColumnCount; j++ {
				horizontalColumn := b.Columns[j]
				horizontalToken := horizontalColumn.GetToken(i)
				switch horizontalToken {
				case "X":
					counter++
					if counter == 4 && token == "X" {
						return Player2
					}
				case "O":
					counter++
					if counter == 4 && token == "O" {
						return Player1
					}
				default:
					counter = 0
				}

			}

		}
	}
	return NoResult
}

// CheckIfGameFinishedVertical checks if someone won vertical
func (b *Board) CheckIfGameFinishedVertical() int {
	var counter int
	for i := 0; i < b.ColumnCount; i++ {
		column := b.Columns[i]
		if column.Used < 4 {
			continue
		}
		for j := 0; j < column.Used; j++ {
			token := column.GetToken(i)
			switch token {
			case "X":
				counter++
				if counter == 4 && token == "X" {
					return Player2
				}
			case "O":
				counter++
				if counter == 4 && token == "O" {
					return Player1
				}
			}
		}
	}
	return NoResult
}
*/

// only checks horizontally right now
// Todo
func (b *Board) CheckIfGameIsFinished(x int, y int, countedType string, counter int) string {
	current := b.GetToken(x, y)
	if current == countedType {
		if current == "." {
			counter = 0

		} else {
			counter++
		}

	} else if b.IsToken(current) {
		counter = 1
	} else {
		counter = 0
	}
	if counter == 4 {
		return current
	}
	if x+1 < b.ColumnCount {
		return b.CheckIfGameIsFinished(x+1, y, current, counter)
	} else {
		return "."
	}
}
