package render

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/johnny-morrice/gosnake/snake/layer"
	"github.com/johnny-morrice/gosnake/snake/tiles"
)

func Render(screen tcell.Screen, layers layer.Layers) {
	for _, layer := range layers {
		for _, tile := range layer.Tiles {
			x := tile.X + layer.OffsetX
			y := tile.Y + layer.OffsetY
			if tile.Type == tiles.Character {
				screen.Put(x, y, string(tile.Rune), tcell.StyleDefault.Bold(true))
			} else {
				text, ok := tileToText[tile.Type]
				if !ok {
					panic(fmt.Sprintf("unknown tile type: %d", tile.Type))
				}
				style := colorToStyle(tile.Color)
				screen.Put(x, y, text, style)
			}

		}
	}
}

func colorToStyle(colorText string) tcell.Style {
	switch colorText {
	case "red":
		return tcell.StyleDefault.Foreground(color.Red)
	case "green":
		return tcell.StyleDefault.Foreground(color.Green)
	case "blue":
		return tcell.StyleDefault.Foreground(color.Blue)
	default:
		return tcell.StyleDefault
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
