package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gdamore/tcell/v3"
)

type App struct {
	screen       tcell.Screen
	logFile      *os.File
	ticker       *time.Ticker
	rootCtx      context.Context
	shutdown     context.CancelFunc
	inputHandler *InputHandler
}

func Setup() (*App, error) {
	logFile, err := initLogging()
	if err != nil {
		return nil, fmt.Errorf("logging: %w", err)
	}

	screen, err := initScreen()
	if err != nil {
		defer func() {
			err := logFile.Close()
			if err != nil {
				slog.Error("failed to close log file", "err", err)
			}
		}()
		return nil, fmt.Errorf("screen: %w", err)
	}

	ticker := time.NewTicker(150 * time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	// TODO: pass a real handler here
	inputHandler := MakeInputHandler(cancel, nil)

	return &App{
		rootCtx:      ctx,
		shutdown:     cancel,
		screen:       screen,
		logFile:      logFile,
		ticker:       ticker,
		inputHandler: inputHandler,
	}, nil
}

func (a *App) Run() {
	defer a.shutdown()
	defer a.ticker.Stop()
	defer a.screen.Fini()
	defer func() {
		if err := a.logFile.Close(); err != nil {
			slog.Error("failed to close log file", "err", err)
		}
	}()

	a.loop()
}

func initLogging() (*os.File, error) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	f, err := os.OpenFile("snake.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	return f, nil
}

func initScreen() (tcell.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}
	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()
	return screen, nil
}

func drawHint(screen tcell.Screen, msg string) {
	_, h := screen.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorGray)
	for i, ch := range msg {
		screen.Put(i, h-1, string(ch), style)
	}
}

func redraw(screen tcell.Screen) {
	screen.Clear()
	drawHint(screen, "q: quit")
	// TODO: draw game state
	screen.Show()
}
