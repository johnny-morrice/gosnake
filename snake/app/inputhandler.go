package app

import (
	"context"
	"log/slog"

	"github.com/gdamore/tcell/v3"
)

type InputHandler struct {
	game     Game
	shutdown context.CancelFunc
}

func MakeInputHandler(shutdown context.CancelFunc, handler Game) *InputHandler {
	return &InputHandler{
		game:     handler,
		shutdown: shutdown,
	}
}

func (h *InputHandler) HandleEventKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyUp:
		h.game.OnPressUp()
	case tcell.KeyDown:
		h.game.OnPressDown()
	case tcell.KeyLeft:
		h.game.OnPressLeft()
	case tcell.KeyRight:
		h.game.OnPressRight()
	default:
		if ev.Str() == "q" {
			slog.Info("shutdown", "reason", "quit key")
			h.shutdown()
		}
	}
}
