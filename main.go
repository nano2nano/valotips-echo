package main

import (
	"os"
	"valotips/controller"
	"valotips/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.GET("/posts", controller.GetPosts)
	e.GET("/posts/:id", controller.GetPost)
	e.POST("/posts", controller.CreatePost)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
