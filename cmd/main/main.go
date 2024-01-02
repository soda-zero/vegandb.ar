package main

import (
	"github.com/labstack/echo/v4"
	"github.com/soda-zero/vegandb/internal/rest"
	"github.com/soda-zero/vegandb/internal/web"
)

func main() {
	e := echo.New()
	web.RegisterRoutes(e)
	rest.RegisterRoutes(e)

	e.Start(":3000")
}
