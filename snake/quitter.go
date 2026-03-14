package snake

import (
	"sync"
)

type Quitter struct {
	Done chan struct{}
	once sync.Once
}

func NewQuitter() *Quitter {
	return &Quitter{
		Done: make(chan struct{}),
	}
}

func (q *Quitter) Quit() {
	q.once.Do(func() { close(q.Done) })
}
