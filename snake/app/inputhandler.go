package app

import (
	"context"
	"log/slog"

	"github.com/gdamore/tcell/v3"
)

type KeyHandler interface {
	OnPressUp()
	OnPressDown()
	OnPressLeft()
	OnPressRight()
}

type InputHandler struct {
	handler  KeyHandler
	shutdown context.CancelFunc
}

func MakeInputHandler(shutdown context.CancelFunc, handler KeyHandler) *InputHandler {
	return &InputHandler{
		handler:  handler,
		shutdown: shutdown,
	}
}

func (h *InputHandler) HandleEventKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyUp:
		h.handler.OnPressUp()
	case tcell.KeyDown:
		h.handler.OnPressDown()
	case tcell.KeyLeft:
		h.handler.OnPressLeft()
	case tcell.KeyRight:
		h.handler.OnPressRight()
	default:
		if ev.Str() == "q" {
			slog.Info("shutdown", "reason", "quit key")
			h.shutdown()
		}
	}
}
