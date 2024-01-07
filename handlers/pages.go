package handlers

import (
    "github.com/yannis94/kagami/views/pages"

	"github.com/labstack/echo/v4"
)

type Pages struct {}

func NewPageHandler() *Pages {
    return &Pages{}
}

func (h Pages) HandleGetIndex(c echo.Context) error {
    return render(c, pages.Index())
}

func (h Pages) HandleGetProjects(c echo.Context) error {
    return render(c, pages.Projects())
}
