package internal

import tm "github.com/buger/goterm"

const (
	Player1 = iota
	Player2
)

type Board struct {
	Slots       []Column // List of columns for the board
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
		Slots:       cs,
		ColumnCount: columns,
		NextPlayer:  Player1,
	}
}

func (b *Board) Print() {
	tm.Clear()
	for x, s := range b.Slots {
		if s.Used == 0 {
			continue
		}
		for i := 0; i < s.Used; i++ {
			tm.MoveCursor(x, s.Capacity-i)
			tm.Print(s.GetToken(i))
		}
		tm.MoveCursor(0, s.Capacity+1)
		tm.Println("123456")
		tm.Flush()
	}
}
