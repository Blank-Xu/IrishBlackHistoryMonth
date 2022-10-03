package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(engine *echo.Echo) {
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	mr := engine.Group("/management")
	_ = mr

}
