package layer

type Layers []Layer

type Layer struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
	Tiles   []Tile
}

type Tile struct {
	X     int
	Y     int
	Type  int
	Rune  rune
	Color string
}
