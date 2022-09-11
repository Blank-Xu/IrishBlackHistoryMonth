package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// echoRequest response request query string for debug
func echoRequest(ctx echo.Context) error {
	return ctx.String(
		http.StatusOK,
		fmt.Sprintf("Path: %s\nQueryString: %s",
			ctx.Path(),
			ctx.QueryString()),
	)
}

// serverTime response server time
func serverTime(ctx echo.Context) error {
	return ctx.String(http.StatusOK, time.Now().Format(time.RFC3339))
}
