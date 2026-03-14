package app

import (
	"log/slog"

	"github.com/gdamore/tcell/v3"
)

func (a *App) loop() {
	redraw(a.screen)
	slog.Info("entering main loop")

	for {
		select {
		case <-a.ticker.C:
			// TODO: update game state
			redraw(a.screen)

		case <-a.rootCtx.Done():
			slog.Info("main loop stopping", "reason", "root context done")
			return

		case ev := <-a.screen.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventKey:
				a.inputHandler.HandleEventKey(ev)

			case *tcell.EventResize:
				a.screen.Sync()
				redraw(a.screen)
			}
		}
	}
}
