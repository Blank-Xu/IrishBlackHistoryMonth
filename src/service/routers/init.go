package routers

import (
	"github.com/labstack/echo/v4"
)

func Register(engine *echo.Echo) {
	engine.GET("/echo", echoRequest)
	engine.POST("/time", serverTime)
}
