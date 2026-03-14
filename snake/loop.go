package snake

import (
	"log/slog"
	"os"
	"time"

	"github.com/gdamore/tcell/v3"
)

func loop(screen tcell.Screen, ticker *time.Ticker, done <-chan os.Signal) {
	redraw(screen)
	slog.Info("entering main loop")

	for {
		select {
		case <-ticker.C:
			// TODO: update game state
			redraw(screen)

		case sig := <-done:
			slog.Info("shutdown", "reason", sig.String())
			return

		case ev := <-screen.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Str() == "q" {
					slog.Info("shutdown", "reason", "quit key")
					return
				}
				// TODO: handle direction keys

			case *tcell.EventResize:
				screen.Sync()
				redraw(screen)
			}
		}
	}
}
