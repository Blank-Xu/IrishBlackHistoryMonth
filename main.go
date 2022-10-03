package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"webservice/cmd"
	"webservice/config"
	managementRouters "webservice/management/routers"
	"webservice/modules/http_server"
	serviceRouters "webservice/service/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("service starting...")

	// using this context for the whole service
	ctx := context.Background()

	// prepare the necessary data
	cmd.Run(ctx)

	cfg := config.GetConfig()

	engine := echo.New()
	managementRouters.Register(engine)
	serviceRouters.Register(engine)

	// create HTTP server
	httpServer := http_server.NewServer(cfg.HTTPServer, engine)
	httpServer.SetKeepAlivesEnabled(true)

	go func(srv *http.Server) {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("http server start failed, err: " + err.Error())
		}
	}(httpServer)

	fmt.Printf("http server started at [%s]\n", httpServer.Addr)

	fmt.Println("please press CTRL+C to quite the process")
	fmt.Println()

	// notify exit signal
	quit := make(chan os.Signal, 1)
	// these signals are for testing purpose only, in formal environment is different.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)

	s := <-quit
	fmt.Println()
	fmt.Printf("received signal: %v\n", s)

	fmt.Println("service will gracefully close in 7 seconds, please wait for a few seconds")
	fmt.Println()

	ctx, cancel := context.WithTimeout(ctx, time.Second*7)
	defer cancel()

	fmt.Println("service stopping ...")

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("http server shutdown get error, err: %s\n", err.Error())
	}

	<-ctx.Done()
	fmt.Println("service stopped")
}
