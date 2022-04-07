package main

import (
	"postgresql-gorm/config"
	"postgresql-gorm/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	config.InitDB()

	e := echo.New()
	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":3990"))

}
