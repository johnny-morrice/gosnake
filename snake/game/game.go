package game

import (
	"errors"
	"math/rand"

	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

type Game struct {
	width      int
	height     int
	geometry   Geometry
	snake      *Snake
	food       *Food
	tickedOnce bool
}

func New(width, height int) (*Game, error) {
	if width < 5 || height < 5 {
		return nil, errors.New("game dimensions must be at least 5x5")
	}

	geometry := &Torus{
		Width:  width - 2,
		Height: height - 3,
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
		food:     NewFood(geometry),
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
	g.snake.Direction = Delta{DX: 1, DY: 0}
}

func (g *Game) Tick() {
	defer func() {
		g.tickedOnce = true
	}()

	const foodChance = 0.01
	if !g.tickedOnce || rand.Float32() < foodChance {
		g.food.AddFood(g.snake.Deque.Seq())
	}
	for point, foodItem := range g.food.Food {
		if g.snake.IsCollide(point) {
			g.food.Eaten(point)
			g.snake.EatFood(foodItem.Nutrition)
		}
	}

	g.snake.Tick()
}

func (g *Game) Render() layer.Layers {
	snakeLayer := g.snake.Render()
	snakeLayer.OffsetX = 1
	snakeLayer.OffsetY = 1
	foodLayer := g.food.Render()
	foodLayer.OffsetX = 1
	foodLayer.OffsetY = 1
	return layer.Layers{
		g.backgroundLayer(),
		foodLayer,
		snakeLayer,
	}
}

func (g *Game) backgroundLayer() layer.Layer {
	// Lines for the game border
	// "q for quit" hint at the bottom
	const msg = "q: quit"
	myTiles := make([]layer.Tile, 0, (g.width*2)+(g.height*2)+len(msg))
	// Corners first
	myTiles = append(myTiles, layer.Tile{X: 0, Y: 0, Type: tiles.Corner})
	myTiles = append(myTiles, layer.Tile{X: g.width - 1, Y: 0, Type: tiles.Corner})
	myTiles = append(myTiles, layer.Tile{X: 0, Y: g.height - 2, Type: tiles.Corner})
	myTiles = append(myTiles, layer.Tile{X: g.width - 1, Y: g.height - 2, Type: tiles.Corner})

	// Lines next so they are under the corners
	for x := 1; x < g.width-1; x++ {
		myTiles = append(myTiles, layer.Tile{X: x, Y: 0, Type: tiles.HorizontalLine})
		myTiles = append(myTiles, layer.Tile{X: x, Y: g.height - 2, Type: tiles.HorizontalLine})
	}
	for y := 1; y < g.height-2; y++ {
		myTiles = append(myTiles, layer.Tile{X: 0, Y: y, Type: tiles.VerticalLine})
		myTiles = append(myTiles, layer.Tile{X: g.width - 1, Y: y, Type: tiles.VerticalLine})
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
