package game

type Game struct {
	Grid    *Grid
	Running bool
}

func NewGame(width, height int) *Game {
	return &Game{
		Grid:    NewGrid(width, height),
		Running: true,
	}
}

func (g *Game) Update() {
	if g.Running {
		g.Grid.NextGeneration()
	}
}
