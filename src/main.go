package main

import (
	"fmt"
	"net/http"
	"time"
	managementRouters "webservice/management/routers"
	serviceRouters "webservice/service/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()

	engine.GET("/echo", echoGetRequest)
	engine.GET("/echo", echoServerTime)

	mr := engine.Group("/management")
	managementRouters.Init(mr)

	sr := engine.Group("/service")
	serviceRouters.Init(sr)

	engine.Start(":9080")
}

func echoGetRequest(ctx echo.Context) error {
	return ctx.String(
		http.StatusOK,
		fmt.Sprintf("Path: %s\nQueryString: %s",
			ctx.Path(),
			ctx.QueryString()),
	)
}

func echoServerTime(ctx echo.Context) error {
	return ctx.String(http.StatusOK, time.Now().Format(time.RFC3339))
}
