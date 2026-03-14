package snake

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

		case sig := <-a.done:
			slog.Info("shutdown", "reason", sig.String())
			return

		case ev := <-a.screen.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Str() == "q" {
					slog.Info("shutdown", "reason", "quit key")
					return
				}
				// TODO: handle direction keys

			case *tcell.EventResize:
				a.screen.Sync()
				redraw(a.screen)
			}
		}
	}
}
