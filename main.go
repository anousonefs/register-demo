package main

import (
	"register/v1/user"
	user2 "register/v2/user"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	user.InitUsers()

	v1 := e.Group("/v1")
	v1.GET("/users", user.ListUsers)
	v1.POST("/users", user.CreateUser)
	v1.PUT("/users/:id", user.UpdateUser)
	v1.DELETE("/users/:id", user.DeleteUser)

	v2 := e.Group("/v2")
	v2.GET("/users", user2.ListUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
