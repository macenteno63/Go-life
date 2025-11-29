package game

import (
	"fmt"
	"math/rand"
)

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

func (g *Grid) SetCell(x, y int, isAlive bool) {
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

func (g *Grid) countAliveNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if g.IsAlive(x+i, y+j) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) NextGeneration() {
	newBoard := make([][]bool, g.Height)
	for y := 0; y < g.Height; y++ {
		newBoard[y] = make([]bool, g.Width)
	}

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			aliveNeighbors := g.countAliveNeighbors(x, y)
			cellAlive := g.IsAlive(x, y)

			if cellAlive && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				newBoard[y][x] = true
			} else if !cellAlive && aliveNeighbors == 3 {
				newBoard[y][x] = true
			} else {
				newBoard[y][x] = false
			}
		}
	}

	g.Board = newBoard
}

func (g *Grid) Clear() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Board[y][x] = false
		}
	}
}

func (g *Grid) Randomize(density float32) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < density {
				g.Board[y][x] = true
			} else {
				g.Board[y][x] = false
			}
		}
	}
}
