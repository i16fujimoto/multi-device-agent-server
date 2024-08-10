package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/multi-device-agent-server/internal/app/ui"
	"github.com/multi-device-agent-server/internal/pkg/cerror"
	"github.com/multi-device-agent-server/internal/pkg/logger"
	"github.com/multi-device-agent-server/internal/pkg/validator"
)

const httpAddr = ":8080"

func main() {
	// env := config.GetEnv()

	// Initialize logger
	logger := logger.New()

	// handler
	handler := ui.NewHandler()

	// server
	e := echo.New()

	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = cerror.CustomHTTPErrorHandler

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[echo] time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	// route
	e.GET("/health", handler.GetHealth)

	// Start server
	go func() {
		if err := e.Start(httpAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server: failed to start HTTP server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) //nolint:gomnd
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logger.Error("server: failed to shutdown server gracefully", zap.Error(err))
	}
}
