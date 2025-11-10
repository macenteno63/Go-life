package main

import "golife/internal/game"

func main() {
	grid := game.NewGrid(5, 5)
	grid.SetSell(0, 1, true)
	grid.Print()
}
