package snake

import (
	"context"
	"log/slog"

	"github.com/gdamore/tcell/v3"
)

type InputHandler struct {
	quitter *Quitter
}

func MakeInputHandler(quitter *Quitter) *InputHandler {
	return &InputHandler{
		quitter: quitter,
	}
}

func (h *InputHandler) HandleEventKey(ctx context.Context, ev *tcell.EventKey) {
	if ev.Str() == "q" {
		slog.Info("shutdown", "reason", "quit key")
		h.quitter.Quit()
	}
}
