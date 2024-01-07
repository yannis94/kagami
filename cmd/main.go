package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yannis94/kagami/handlers"
)

func main() {
    pageHandler := handlers.NewPageHandler()
    hypermediaHandler := handlers.NewHypermediaHandler()
    server := echo.New()

    server.Static("/static", "src")
    hm := server.Group("/hm")

    server.GET("/", pageHandler.HandleGetIndex)
    server.GET("/projects", pageHandler.HandleGetIndex)
    
    hm.GET("/contact", hypermediaHandler.HandleGetContact)

    server.Logger.Fatal(server.Start(":3000"))
}
