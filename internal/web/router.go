package web

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	homeHandler := &HomeHandler{}
	e.GET("/", homeHandler.HandleHome)
	e.GET("/search", homeHandler.HandleListProducts)
}
