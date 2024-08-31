package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"app/pkg"
	"github.com/alexflint/go-arg"
	"github.com/skovtunenko/graterm"
)

var version = "will-come-from-ld-flags"

type args struct {
	GracePeriod time.Duration `arg:"--grace-period, -g, env:GRACE_PERIOD" default:"5s"    help:"grace period for termination" placeholder:"duration"`
	Verbose     bool          `arg:"-v, env"                              default:"false" help:"print verbose output"`
}

func (args) Version() string {
	return version
}

func (args) Description() string {
	return "Sample CLI to use as a template"
}

func main() {
	var a args
	p := arg.MustParse(&a)
	out := func(msg string) {
		if _, err := fmt.Fprintln(os.Stdout, msg); err != nil {
			p.Fail(err.Error())
		}
	}
	app := pkg.NewApp(a.Verbose, out)

	cancelableCtx, cancel := context.WithCancel(context.Background())

	terminator, ctx := graterm.NewWithSignals(cancelableCtx, syscall.SIGINT, syscall.SIGTERM)
	terminator.WithOrder(0).Register(a.GracePeriod, app.Close)

	go func() {
		if err := app.Run(ctx); err != nil {
			p.Fail(err.Error())
		}
		cancel()
	}()

	if err := terminator.Wait(ctx, a.GracePeriod); err != nil {
		p.Fail(err.Error())
	}
}
