package view

import (
	"golife/internal/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	CellSize int
}

func NewRenderer(cellSize int) *Renderer {
	return &Renderer{CellSize: cellSize}
}

func (r *Renderer) Render(grid *game.Grid) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.IsAlive(x, y) {
				rl.DrawRectangle(int32(x*r.CellSize), int32(y*r.CellSize), int32(r.CellSize), int32(r.CellSize), rl.White)
			} else {
				rl.DrawRectangleLines(int32(x*r.CellSize), int32(y*r.CellSize), int32(r.CellSize), int32(r.CellSize), rl.DarkGray)
			}
		}
	}

	rl.EndDrawing()
}
