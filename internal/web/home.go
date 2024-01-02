package web

import (
	"github.com/labstack/echo/v4"
	"github.com/soda-zero/vegandb/www/home"
)

func HomeHandler(c echo.Context) error {
	return render(c, home.App())
}
