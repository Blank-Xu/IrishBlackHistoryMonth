package routers

import (
	"github.com/labstack/echo/v4"
)

func Init(engine *echo.Echo) {
	engine.GET("/echo", echoRequest)
	engine.GET("/time", serverTime)
}
