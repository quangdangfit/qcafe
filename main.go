package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"qcafe/routers"
	"time"
)

func main() {
	e := echo.New()

	routers.SetRouters(e)

	// Start server
	go func() {
		port := "8080"
		log.Info("Starting at port: " + port)
		err := e.Start(":" + port)
		if err != nil {
			log.Error(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
