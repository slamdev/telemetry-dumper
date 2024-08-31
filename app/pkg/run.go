package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type App struct {
	verbose bool
	srv     *http.Server
	out     func(msg string)
}

func NewApp(verbose bool, out func(msg string), httpPort int) *App {
	app := App{verbose: verbose, out: out}
	app.srv = &http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		Handler:      app.createHTTPHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	slog.Info("server is created", "port", httpPort)
	return &app
}

func (a *App) createHTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	slog.Info("health handler is created", "path", "/health")
	mux.HandleFunc("/api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("failed to read body", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		headers, err := json.Marshal(r.Header)
		if err != nil {
			slog.Error("failed to marshal headers", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		a.out(fmt.Sprintf("req: %s, body: %s, headers: %s", r.RequestURI, body, headers))
		w.WriteHeader(http.StatusOK)
	})
	slog.Info("prometheus remote write handler is created", "path", "/api/v1/push")
	return mux
}

func (a *App) Run(ctx context.Context) error {
	a.srv.BaseContext = func(net.Listener) context.Context { return ctx }
	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}

func (a *App) Close(ctx context.Context) {
	slog.InfoContext(ctx, "closing http server")
	if err := a.srv.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to shutdown http server", "err", err)
	}
}
