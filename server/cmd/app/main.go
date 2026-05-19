package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/logger"
	core_http_middleware "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/transport/http/middleware"
	core_http_server "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/transport/server"
	users_transport_http "95.174.104.37/gitlab/vifrolov/todo_app/internal/features/users/transport/http"

	"go.uber.org/zap"
)

func main() {
	ctx, cansel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cansel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}

	defer logger.Close()

	logger.Debug("started ToDo application")

	usersTransportHTTP := users_transport_http.NewUsersHttpHandler(nil)
	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoute(usersRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.ConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error: %w", zap.Error(err))
	}

}
