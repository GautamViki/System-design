package game

import "fmt"

type Board struct {
	Grid [][]rune
	Size int
}

func NewBoard(size int) *Board {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}
	return &Board{
		Grid: grid,
		Size: size,
	}
}

func (b *Board) PrintBoard() {
	for i := range b.Grid {
		for j := range b.Grid[i] {
			if b.Grid[i][j] == 0 {
				fmt.Print("_")
			} else {
				fmt.Printf("%c", b.Grid[i][j])
			}
			if j < b.Size-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) IsMoveValid(x, y int) bool {
	return x >= 0 && x < b.Size && y >= 0 && y < b.Size && b.Grid[x][y] == 0
}

func (b *Board) PlaceMark(x, y int, mark rune) {
	if b.IsMoveValid(x, y) {
		b.Grid[x][y] = mark
	}
}

func (b *Board) IsGameOver() bool {
	return b.HasWinner() || b.IsFull()
}

func (b *Board) HasWinner() bool {
	// Logic to check rows, columns, and diagonals for a winner
	return false
}

func (b *Board) IsFull() bool {
	for _, row := range b.Grid {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}
