package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers"
	"github.com/PingTeeratorn789/reverse_proxy/internal/handlers/middlewares"
	"github.com/PingTeeratorn789/reverse_proxy/server/router/routers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	serverShutdown sync.WaitGroup
	loggerr        *zap.Logger
	serverHeader   = "Fiber"
)

func Start() (err error) {

	loggerr = app.logger.GetLogger()
	defer app.logger.Close()
	server := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  serverHeader,
		AppName:       app.confs.App.Name,
		BodyLimit:     1024 * 1024 * 102, //102MB
		ErrorHandler:  middlewares.ErrorHandler(app.confs.App.Prefix),
	})
	// After Shutdown
	server.Hooks().OnShutdown(AfterShutdownServer)
	serverShutdown.Add(1)
	handlers := handlers.NewHandlers(handlers.Dependencies{
		Config: *app.confs,
	})

	routers.SetupRouters(server, *handlers)
	// Wait Signal Ctrl+C or SIGTERM
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer serverShutdown.Done()
		<-gracefulStop
		loggerr.Info("Received shutdown signal. Gracefully stopping...")
		shutdownServer(server)
	}()

	portHTTP := fmt.Sprintf(":%s", app.confs.App.Port)
	if err := server.Listen(portHTTP); err != nil {
		loggerr.Error("Error starting server", zap.Error(err))
		log.Fatal(err)
	}

	serverShutdown.Wait()
	close(gracefulStop)
	loggerr.Info("Server has been closed gracefully")
	return err
}

func shutdownServer(server *fiber.App) {
	if err := server.Shutdown(); err != nil {
		loggerr.Error("Error shutting down server", zap.Error(err))
		log.Fatal(err)
	}
}

func AfterShutdownServer() error {
	loggerr.Info("Server Shutdown Completed")
	return nil
}
