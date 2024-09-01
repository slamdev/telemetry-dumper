package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/grafana/loki/pkg/logproto"
	"github.com/xhhuango/json"

	"github.com/klauspost/compress/snappy"
	"github.com/prometheus/prometheus/prompb"
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
	mux.HandleFunc("/prom/api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		if err := a.handlePrometheusRequest(r); err != nil {
			slog.ErrorContext(r.Context(), "failed to handle prometheus request", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	slog.Info("prometheus remote write handler is created", "path", "/prom/api/v1/push")
	mux.HandleFunc("/loki/api/v1/push", func(w http.ResponseWriter, r *http.Request) {
		if err := a.handleLokiRequest(r); err != nil {
			slog.ErrorContext(r.Context(), "failed to handle loki request", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	slog.Info("loki handler is created", "path", "/loki/api/v1/push")
	return mux
}

func (a *App) handleLokiRequest(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	decompressedBytes, err := snappy.Decode(nil, body)
	if err != nil {
		return fmt.Errorf("failed to decode body: %w", err)
	}
	lokiReq := logproto.PushRequest{}
	if err := lokiReq.Unmarshal(decompressedBytes); err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}
	if err := a.log(lokiReq, r); err != nil {
		return fmt.Errorf("failed to log request: %w", err)
	}
	return nil
}

func (a *App) handlePrometheusRequest(r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	decompressedBytes, err := snappy.Decode(nil, body)
	if err != nil {
		return fmt.Errorf("failed to decode body: %w", err)
	}
	promReq := prompb.WriteRequest{}
	if err := promReq.Unmarshal(decompressedBytes); err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}
	if err := a.log(promReq, r); err != nil {
		return fmt.Errorf("failed to log request: %w", err)
	}
	return nil
}

func (a *App) log(body any, r *http.Request) error {
	out := map[string]any{}
	out["body"] = body
	out["headers"] = r.Header
	out["req"] = r.RequestURI
	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}
	a.out(string(b))
	return nil
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
