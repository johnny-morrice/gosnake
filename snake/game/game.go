package game

type Game struct {
	geometry Geometry
	snake    *Snake
}

func New() *Game {
	geometry := &Torus{
		Width:  width,
		Height: height,
	}
	snake := NewSnake(
		Point{X: geometry.Width / 2, Y: geometry.Height / 2},
		Delta{DX: 1, DY: 0},
		geometry)
	return &Game{
		geometry: geometry,
		snake:    snake,
	}
}

func (g *Game) OnPressUp() {
	g.snake.Direction = Delta{DX: 0, DY: -1}
}

func (g *Game) OnPressDown() {
	g.snake.Direction = Delta{DX: 0, DY: 1}
}

func (g *Game) OnPressLeft() {
	g.snake.Direction = Delta{DX: -1, DY: 0}
}

func (g *Game) OnPressRight() {
	g.snake.Direction = Delta{DX: 0, DY: 1}
}

func (g *Game) Tick() {
	g.snake.Tick()
}

const width = 20
const height = 20
