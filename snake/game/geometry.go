package game

import (
	"iter"
	"math/rand"
)

type Point struct {
	X int
	Y int
}

type Delta struct {
	DX int
	DY int
}

type Geometry interface {
	Add(p Point, d Delta) Point
	RandomPoint(occupied ...iter.Seq[Point]) Point
}

type Torus struct {
	Width  int
	Height int
}

func (t Torus) Add(p Point, d Delta) Point {
	x := (p.X + d.DX) % t.Width
	y := (p.Y + d.DY) % t.Height

	if x < 0 {
		x += t.Width
	}

	if y < 0 {
		y += t.Height
	}

	return Point{X: x, Y: y}
}

func (t Torus) RandomPoint(occupied ...iter.Seq[Point]) Point {
	occupiedMap := make(map[Point]struct{})
	for _, mySeq := range occupied {
		for p := range mySeq {
			occupiedMap[p] = struct{}{}
		}
	}

	total := t.Width * t.Height
	free := total - len(occupiedMap)
	if free == 0 {
		panic("no free cells")
	}

	n := rand.Intn(free)
	for i := range total {
		p := Point{X: i % t.Width, Y: i / t.Width}
		_, isOccupied := occupiedMap[p]
		if !isOccupied {
			if n == 0 {
				return p
			}
			n--
		}
	}
	panic("unreachable")
}
