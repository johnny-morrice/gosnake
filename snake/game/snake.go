package game

import (
	"github.com/johnny-morrice/gosnake/snake/deque"
	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tile"
)

type Snake struct {
	Deque     deque.Deque[Point]
	Direction Delta
	Food      int
	Geometry  Geometry
}

func NewSnake(initialPosition Point, initialDirection Delta, geometry Geometry) *Snake {
	d := deque.New[Point]()
	d.PushFront(initialPosition)

	return &Snake{
		Deque:     *d,
		Direction: initialDirection,
		Geometry:  geometry,
	}
}

func (s *Snake) Tick() {
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
	tiles := make([]layer.Tile, len(s.Deque))
	maxX := 0
	maxY := 0
	for i, point := range s.Deque {
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
		tileType := tile.SnakeBody
		if i == 0 {
			tileType = tile.SnakeHead
		}
		tiles[i] = layer.Tile{
			X:    point.X,
			Y:    point.Y,
			Type: tileType,
		}
	}
	return layer.Layer{
		Width:  maxX + 1,
		Height: maxY + 1,
		Tiles:  tiles,
	}
}
