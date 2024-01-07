package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/yannis94/kagami/views/components/popup"
)

type Hypermedia struct {}

func NewHypermediaHandler() *Hypermedia {
    return &Hypermedia{}
}

func (h Hypermedia) HandleGetContact(c echo.Context) error {
    return render(c, popup.Contact())
}
