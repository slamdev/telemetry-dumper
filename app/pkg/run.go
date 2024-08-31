package pkg

import (
	"context"
)

type App struct {
	verbose bool
	out     func(msg string)
}

func NewApp(verbose bool, out func(msg string)) *App {
	return &App{verbose: verbose, out: out}
}

func (a *App) Run(_ context.Context) error {
	if a.verbose {
		a.out("running in verbose mode")
	} else {
		a.out("running")
	}
	return nil
}

func (a *App) Close(_ context.Context) {
}
