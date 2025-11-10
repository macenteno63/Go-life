package game

import "fmt"

type Grid struct {
	Width  int
	Height int
	Board  [][]bool
}

func NewGrid(width, height int) *Grid {
	board := make([][]bool, height)

	for i := range height {
		board[i] = make([]bool, width)
	}

	return &Grid{
		Width:  width,
		Height: height,
		Board:  board,
	}
}

func (g *Grid) inBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) SetSell(x, y int, isAlive bool) {
	if g.inBounds(x, y) {
		g.Board[y][x] = isAlive
	}
}

func (g *Grid) IsAlive(x, y int) bool {
	if g.inBounds(x, y) {
		return g.Board[y][x]
	}
	return false
}

func (g *Grid) Print() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.IsAlive(x, y) {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
