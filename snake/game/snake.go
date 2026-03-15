package game

import (
	"slices"

	"github.com/johnny-morrice/gosnake/snake/deque"
	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

type Snake struct {
	Deque          deque.Deque[Point]
	Direction      Delta
	Food           int
	Geometry       Geometry
	ticksSinceMove int
	speed          int
}

func NewSnake(initialPosition Point, initialDirection Delta, geometry Geometry) *Snake {
	d := deque.New[Point]()
	d.PushFront(initialPosition)

	return &Snake{
		Deque:          *d,
		Direction:      initialDirection,
		Geometry:       geometry,
		speed:          initialSpeed,
		ticksSinceMove: 0,
	}
}

func (s *Snake) EatFood(nutrition int) {
	s.Food += nutrition
	if s.speed > 0 {
		s.speed--
	}
}

func (s *Snake) IsCollide(point Point) bool {
	return slices.Contains(s.Deque, point)
}

func (s *Snake) Tick() {
	if s.ticksSinceMove < s.speed {
		s.ticksSinceMove++
		return
	}
	s.ticksSinceMove = 0

	head := s.Deque[0]
	newHead := s.Geometry.Add(head, s.Direction)
	s.Deque.PushFront(newHead)

	if s.Food > 0 {
		s.Food--
	} else {
		s.Deque.PopBack()
	}
}

func (s *Snake) Render() layer.Layer {
	myTiles := make([]layer.Tile, len(s.Deque))
	maxX := 0
	maxY := 0
	for i, point := range s.Deque {
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
		tileType := tiles.SnakeBody
		if i == 0 {
			tileType = tiles.SnakeHead
		}
		myTiles[i] = layer.Tile{
			X:     point.X,
			Y:     point.Y,
			Type:  tileType,
			Color: "green",
		}
	}
	return layer.Layer{
		Width:  maxX + 1,
		Height: maxY + 1,
		Tiles:  myTiles,
	}
}

const initialSpeed = 10
