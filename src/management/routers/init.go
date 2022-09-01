package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(engine *echo.Group) {
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())
}
