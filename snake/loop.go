package snake

import (
	"context"
	"log/slog"

	"github.com/gdamore/tcell/v3"
)

func (a *App) loop(ctx context.Context) {
	redraw(a.screen)
	slog.Info("entering main loop")

	for {
		select {
		case <-a.ticker.C:
			// TODO: update game state
			redraw(a.screen)

		case <-a.quitter.Done:
			slog.Info("main loop stopping", "reason", "quitter channel triggered")
			return

		case ev := <-a.screen.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventKey:
				a.inputHandler.HandleEventKey(ctx, ev)

			case *tcell.EventResize:
				a.screen.Sync()
				redraw(a.screen)
			}
		}
	}
}
