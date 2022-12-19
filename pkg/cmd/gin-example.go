package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/IkeMurami-Examples/go-gin-example/pkg/server"
	"github.com/IkeMurami-Examples/go-gin-example/pkg/utils"
)

const (
	serverShutdownTimeout = 5
)

var logger *zap.Logger

func StartServer(ctx context.Context) error {
	logger = utils.LoggerFromContext(ctx)
	// Get Context with cancel
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Setting interrupt
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// Get context for running Go-routins
	group, ctx := errgroup.WithContext(ctx)

	// Run the Gin server
	httpServer, err := startHTTPServer(ctx, group)
	if err != nil {
		return fmt.Errorf("Failed to start http server: %v", err)
	}

	// Catch the shutdown signal
	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}
	logger.Warn("received shutdown signal")

	// Shutdown server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(),
		serverShutdownTimeout*time.Second)
	defer shutdownCancel()

	_ = httpServer.Shutdown(shutdownCtx)

	return group.Wait()
}

func startHTTPServer(ctx context.Context, group *errgroup.Group) (*http.Server, error) {
	endpoint := viper.GetString("HTTP_ENDPOINT")

	mux := http.NewServeMux()

	srv := server.NewServer(ctx, endpoint, mux)

	group.Go(func() error {
		logger.Info("HTTP Server serving", zap.String("endpoint", endpoint))

		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("HTTP Server stopped", zap.String("endpoint", endpoint), zap.Error(err))
		}

		return err
	})

	return srv, nil
}
