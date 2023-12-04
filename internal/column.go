package internal

import (
	"fmt"
	"os"
)

const (
	Player1Token = 1
	Player2Token = 98
)

type Column struct {
	Slot     []int // This array holds the token
	Capacity int   // How many token can be in the slot
	Used     int   // How many token are currently in the slot
}

func NewColumn(capacity int) *Column {
	return &Column{
		Slot:     make([]int, capacity),
		Capacity: capacity,
		Used:     0,
	}
}

// IsFull returns true if column is filled
func (c *Column) IsFull() bool {
	if c.Used == c.Capacity {
		return true
	}
	return false
}

func (c *Column) GetToken(slot int) string {
	switch c.Slot[slot] {
	case Player1Token:
		return "O"
	case Player2Token:
		return "X"
	}
	return "."
}

func (c *Column) PlaceToken(playerToken int) {
	if c.Used == c.Capacity {
		fmt.Println("Error: Slot is filled")
		os.Exit(1)
	}
	c.Slot[c.Used] = playerToken
	c.Used++
}
