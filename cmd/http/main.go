package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yannis94/kagami/handlers"
	"github.com/yannis94/kagami/middlewares"
	"github.com/yannis94/kagami/store"
)

func init() {
    err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
    db := store.NewSQLiteStorage()
    db.Init()
    
    pageHandler := handlers.NewPageHandler(db)
    hypermediaHandler := handlers.NewHypermediaHandler(db)
    adminHandler := handlers.NewAdminHandler(db)

    server := echo.New()

    server.Static("/static", "src")

    server.Use(middleware.Logger())
    server.Use(middleware.Recover())

    server.GET("/", pageHandler.HandleGetIndex)
    server.GET("/projects", pageHandler.HandleGetProjects)
    server.GET("/projects/:id", pageHandler.HandleGetProject)
    
    hm := server.Group("/hm")
    hm.GET("/contact", hypermediaHandler.HandleGetContact)
    hm.GET("/skills", hypermediaHandler.HandleGetSkills)
    hm.GET("/skills/categories", hypermediaHandler.HandleGetSkillCategories)
    hm.GET("/projects/keywords", hypermediaHandler.HandleGetProjectKeywords)
    hm.POST("/projects", hypermediaHandler.HandleGetProjects)
    hm.POST("/yayadmin/login", adminHandler.HandlePostLogin)

    hmAdmin := hm.Group("/yayadmin")
    hmAdmin.Use(echojwt.WithConfig(middlewares.AuthMiddlewareConfig))
    hmAdmin.GET("/skills", hypermediaHandler.HandleGetAdminSkills)
    hmAdmin.GET("/projects", hypermediaHandler.HandleGetAdminProjects)
    hmAdmin.PUT("/projects/:id", adminHandler.HandlePutProject)
    hmAdmin.PUT("/skills/:id", adminHandler.HandlePutSkill)
    hmAdmin.DELETE("/skills/:id", adminHandler.HandleDeleteSkill)
    hmAdmin.DELETE("/projects/:id", adminHandler.HandleDeleteProject)

    admin := server.Group("/yayadmin")
    admin.GET("/login", pageHandler.HandleGetAdminLogin)
    admin.Use(echojwt.WithConfig(middlewares.AuthMiddlewareConfig))
    admin.GET("", pageHandler.HandleGetAdminHome)
    admin.GET("/projects/:id", pageHandler.HandleGetAdminProject)
    admin.GET("/skills/:id", pageHandler.HandleGetAdminSkill)
    admin.POST("/skill", adminHandler.HandlePostSkill)
    admin.POST("/project", adminHandler.HandlePostProject)

    server.Logger.Fatal(server.Start(os.Getenv("PORT")))
}
