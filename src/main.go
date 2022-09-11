package main

import (
	managementRouters "webservice/management/routers"
	serviceRouters "webservice/service/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	engine := echo.New()

	managementRouters.Init(engine)

	serviceRouters.Init(engine)

	engine.Start(":9080")
}
