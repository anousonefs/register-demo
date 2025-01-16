package main

import (
	"database/sql"
	"fmt"
	"os"
	"register/v1/user"
	user2 "register/v2/user"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	connStr := GetEnv("DB_URL", "")
	fmt.Printf("db_url: %v\n", connStr)

	connStr = "postgresql://postgres:irBNRUdFqGyqZGcjBKPvNFmCbFdzGZcT@viaduct.proxy.rlwy.net:51792/postgres"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Printf("connect db success!\n")

	e := echo.New()
	user.InitUsers()

	v1 := e.Group("/v1")
	v1.GET("/users", user.ListUsers)
	v1.POST("/users", user.CreateUser)
	v1.PUT("/users/:id", user.UpdateUser)
	v1.DELETE("/users/:id", user.DeleteUser)

	// create user service 2
	userService2 := user2.NewUserService(db)

	v2 := e.Group("/v2")
	v2.GET("/users", userService2.ListUsers)
	v2.POST("/users", userService2.CreateUser)
	v2.PUT("/users/:id", userService2.UpdateUser)
	v2.DELETE("/users/:id", userService2.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
