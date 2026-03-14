package game

import "github.com/johnny-morrice/gosnake/snake/deque"

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
