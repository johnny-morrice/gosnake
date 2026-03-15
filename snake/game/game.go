package game

import (
	"errors"

	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

type Game struct {
	width    int
	height   int
	geometry Geometry
	snake    *Snake
}

func New(width, height int) (*Game, error) {
	if width < 5 || height < 5 {
		return nil, errors.New("game dimensions must be at least 5x5")
	}

	geometry := &Torus{
		Width:  width - 2,
		Height: height - 2,
	}
	snake := NewSnake(
		Point{X: geometry.Width / 2, Y: geometry.Height / 2},
		Delta{DX: 1, DY: 0},
		geometry)

	game := &Game{
		width:    width,
		height:   height,
		geometry: geometry,
		snake:    snake,
	}
	return game, nil
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

func (g *Game) Render() layer.Layers {
	snakeLayer := g.snake.Render()
	snakeLayer.OffsetX = 1
	snakeLayer.OffsetY = 1
	return layer.Layers{
		g.backgroundLayer(),
		snakeLayer,
	}
}

func (g *Game) backgroundLayer() layer.Layer {
	// Lines for the game border
	// "q for quit" hint at the bottom
	const msg = "q: quit"
	myTiles := make([]layer.Tile, 0, (g.width*2)+(g.height*2)+len(msg))
	for x := range g.width {
		myTiles = append(myTiles, layer.Tile{X: x, Y: 0, Type: tiles.HorizontalLine})
		myTiles = append(myTiles, layer.Tile{X: x, Y: g.height - 2, Type: tiles.HorizontalLine})
	}
	for y := range g.height {
		myTiles = append(myTiles, layer.Tile{X: 0, Y: y, Type: tiles.VerticalLine})
		myTiles = append(myTiles, layer.Tile{X: g.width - 2, Y: y, Type: tiles.VerticalLine})
	}
	for i, ch := range msg {
		myTiles = append(myTiles, layer.Tile{X: i, Y: g.height - 1, Type: tiles.Character, Rune: ch})
	}
	return layer.Layer{
		Width:  g.width,
		Height: g.height,
		Tiles:  myTiles,
	}
}
