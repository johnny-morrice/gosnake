package render

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

func Render(screen tcell.Screen, layers layer.Layers) {
	for _, layer := range layers {
		for _, tile := range layer.Tiles {
			x := tile.X + layer.OffsetX
			y := tile.Y + layer.OffsetY
			if tile.Type == tiles.Character {
				screen.Put(x, y, string(tile.Rune), tcell.StyleDefault)
			} else {
				text, ok := tileToText[tile.Type]
				if !ok {
					panic(fmt.Sprintf("unknown tile type: %d", tile.Type))
				}
				screen.Put(x, y, text, tcell.StyleDefault)
			}

		}
	}
}

var tileToText = map[int]string{
	tiles.VerticalLine:   "│",
	tiles.HorizontalLine: "─",
	tiles.SnakeHead:      "#",
	tiles.SnakeBody:      "#",
	tiles.SmallFood:      "o",
	tiles.LargeFood:      "O",
	tiles.Corner:         "+",
}
