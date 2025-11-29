package main

import (
	"time"

	"golife/internal/game"
	"golife/internal/view"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width := 100
	height := 100
	cellSize := 8
	updateDelay := 100 * time.Millisecond

	// Init Raylib
	rl.InitWindow(int32(width*cellSize), int32(height*cellSize), "Game of Life")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	// Init Game (remplace grid := ...)
	g := game.NewGame(width, height)
	// madePulsar(g.Grid)

	// Init Renderer
	renderer := view.NewRenderer(cellSize)

	lastUpdate := time.Now()

	for !rl.WindowShouldClose() {
		// Handle Even TODO: put in a function
		if rl.IsKeyPressed(rl.KeySpace) {
			g.Running = !g.Running
		}

		if rl.IsKeyPressed(rl.KeyR) {
			g.Grid.Randomize(0.2)
		}

		if rl.IsKeyPressed(rl.KeyC) {
			g.Grid.Clear()
		}

		mousePosition := rl.GetMousePosition()
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			x := int(mousePosition.X) / cellSize
			y := int(mousePosition.Y) / cellSize
			g.Grid.SetCell(x, y, true)
		}

		if rl.IsMouseButtonPressed(rl.MouseRightButton) {
			x := int(mousePosition.X) / cellSize
			y := int(mousePosition.Y) / cellSize
			g.Grid.SetCell(x, y, false)
		}

		// Game logique
		if g.Running && time.Since(lastUpdate) >= updateDelay {
			g.Update()
			lastUpdate = time.Now()
		}

		// Render game
		renderer.Render(g.Grid)
	}
}

func madePulsar(grid *game.Grid) {
	offsetX := 30
	offsetY := 30

	coords := [][2]int{
		// Haut centre
		{2, 0}, {3, 0}, {4, 0}, {8, 0}, {9, 0}, {10, 0},
		{2, 1}, {3, 1}, {4, 1}, {8, 1}, {9, 1}, {10, 1},
		{2, 2}, {3, 2}, {4, 2}, {8, 2}, {9, 2}, {10, 2},

		// Milieu haut gauche
		{0, 2}, {0, 3}, {0, 4},
		{5, 2}, {5, 3}, {5, 4},

		// Milieu haut droite
		{7, 2}, {7, 3}, {7, 4},
		{12, 2}, {12, 3}, {12, 4},

		// Milieu bas gauche
		{0, 7}, {0, 8}, {0, 9},
		{5, 7}, {5, 8}, {5, 9},

		// Milieu bas droite
		{7, 7}, {7, 8}, {7, 9},
		{12, 7}, {12, 8}, {12, 9},

		// Bas centre
		{2, 12}, {3, 12}, {4, 12}, {8, 12}, {9, 12}, {10, 12},
		{2, 13}, {3, 13}, {4, 13}, {8, 13}, {9, 13}, {10, 13},
		{2, 14}, {3, 14}, {4, 14}, {8, 14}, {9, 14}, {10, 14},
	}

	for _, c := range coords {
		grid.SetCell(offsetX+c[0], offsetY+c[1], true)
	}
}
