package app

import "github.com/johnny-morrice/gosnake/snake/layer"

type Game interface {
	OnPressUp()
	OnPressDown()
	OnPressLeft()
	OnPressRight()
	Tick()
	Render() layer.Layers
}
