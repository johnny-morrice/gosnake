package app

import "github.com/johnny-morrice/gosnake/snake/render"

func (a *App) redraw() {
	a.screen.Clear()

	layers := a.game.Render()
	render.Render(a.screen, layers)

	a.screen.Show()
}
