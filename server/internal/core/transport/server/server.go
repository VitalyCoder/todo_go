package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/logger"

	core_http_middleware "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/transport/http/middleware"

	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (h *HTTPServer) RegisterAPIRoutes(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	h.log.Warn("start HTTP server", zap.String("addr", h.config.Addr))

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		{
			if err != nil {
				return fmt.Errorf("listen anf server HTTP %w", err)
			}
		}
	case <-ctx.Done():
		{
			h.log.Warn("shutdown HTTP server...")
			shutdownCtx, cansel := context.WithTimeout(
				context.Background(),
				h.config.ShutdownTimeout,
			)
			defer cansel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				_ = server.Close()

				return fmt.Errorf("shutdown HTTP server: %w", err)
			}

			h.log.Warn("HTTP server stoped")
		}
	}
	return nil
}
