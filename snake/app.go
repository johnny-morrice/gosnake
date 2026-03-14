package snake

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gdamore/tcell/v3"
)

type App struct {
	screen  tcell.Screen
	logFile *os.File
	cancel  context.CancelFunc
	ticker  *time.Ticker
	done    chan os.Signal
}

func Setup() (*App, error) {
	logFile, err := initLogging()
	if err != nil {
		return nil, fmt.Errorf("logging: %w", err)
	}

	screen, err := initScreen()
	if err != nil {
		return nil, fmt.Errorf("screen: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(150 * time.Millisecond)
	done := make(chan os.Signal, 1)

	startSignalHandler(ctx, done)

	return &App{
		screen:  screen,
		logFile: logFile,
		cancel:  cancel,
		ticker:  ticker,
		done:    done,
	}, nil
}

func (a *App) Run() {
	defer a.cancel()
	defer a.ticker.Stop()
	defer a.screen.Fini()
	defer func() {
		if err := a.logFile.Close(); err != nil {
			slog.Error("failed to close log file", "err", err)
		}
	}()

	loop(a.screen, a.ticker, a.done)
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

// --- Signal handler -------------------------------------------------------

func startSignalHandler(ctx context.Context, done chan<- os.Signal) {
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigCh:
			slog.Info("received signal", "signal", sig)
			done <- sig
		case <-ctx.Done():
		}
	}()
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
