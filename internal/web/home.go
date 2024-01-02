package web

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/soda-zero/vegandb/internal/database"
	"github.com/soda-zero/vegandb/internal/products"
	"github.com/soda-zero/vegandb/www/home"
	"net/http"
)

const (
	hxRequestHeaderKey   = "HX-Request"
	hxRequestHeaderValue = "true"
	defaultRedirectPath  = "/"
)

type HomeHandler struct {
	filter           string
	filteredProducts []database.Product
}

func (h *HomeHandler) HandleHome(c echo.Context) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	productsRepo := products.NewProductRepository(database.New(db))
	h.filter = c.QueryParam("filter")

	h.filteredProducts, err = productsRepo.FilterProducts(context.Background(), h.filter)
	if err != nil {
		return err
	}

	return render(c, home.App(h.filter, h.filteredProducts))
}

func (h *HomeHandler) HandleListProducts(c echo.Context) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	h.filter = c.QueryParam("filter")
	hxRequestHeader := c.Request().Header.Get(hxRequestHeaderKey)

	productsRepo := products.NewProductRepository(database.New(db))
	h.filteredProducts, err = productsRepo.FilterProducts(context.Background(), h.filter)
	if err != nil {
		return err
	}

	if hxRequestHeader == hxRequestHeaderValue {
		if h.filter == "" {
			c.Response().Header().Set("HX-Push-Url", "/")
		}
		return render(c, home.SearchResults(h.filter, h.filteredProducts))

	} else {
		if h.filter == "" && c.Path() == "/search" {
			c.Redirect(http.StatusTemporaryRedirect, defaultRedirectPath)
		}
		return render(c, home.App(h.filter, h.filteredProducts))
	}
}
