package game

import (
	"errors"

	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tile"
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
	return layer.Layers{
		g.backgroundLayer(),
		g.snake.Render(),
	}
}

func (g *Game) backgroundLayer() layer.Layer {
	// Lines for the game border
	// "q for quit" hint at the bottom
	const msg = "q: quit"
	tiles := make([]layer.Tile, 0, (g.width*2)+(g.height*2)+len(msg))
	for x := range g.width {
		tiles = append(tiles, layer.Tile{X: x, Y: 0, Type: tile.HorizontalLine})
		tiles = append(tiles, layer.Tile{X: x, Y: g.height - 2, Type: tile.HorizontalLine})
	}
	for y := range g.height {
		tiles = append(tiles, layer.Tile{X: 0, Y: y, Type: tile.VerticalLine})
		tiles = append(tiles, layer.Tile{X: g.width - 2, Y: y, Type: tile.VerticalLine})
	}
	for i, ch := range msg {
		tiles = append(tiles, layer.Tile{X: i, Y: g.height - 1, Type: tile.Character, Character: ch})
	}
	return layer.Layer{
		Width:  g.width,
		Height: g.height,
		Tiles:  tiles,
	}
}
