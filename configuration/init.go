package configuration

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type configuration struct {
	Server     *echo.Echo
	portServer string
	DB         *sql.DB
}

func Init() *configuration {
	var configuration configuration

	return configuration.
		initTimezone().
		initPostgreSql().
		migrate().
		initServer().
		initService()
}

func (c *configuration) Start() {
	go func() {
		defer func() {
			if err, ok := recover().(error); ok && err != nil {
				log.Printf("[configuration][Start] recover error: %v\n", err)
			}
		}()

		if err := c.Server.Start(":" + c.portServer); err != http.ErrServerClosed {
			log.Fatalf("[configuration][Start] shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func (c *configuration) Stop() {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := c.Server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("[configuration][Stop] shutting down serverr %v\n", err)
	}
}
