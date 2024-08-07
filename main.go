package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// database
var users []User

var id int

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func initUsers() {
	id++
	item := User{
		ID:        id,
		FirstName: "jojo",
		LastName:  "jj",
		Gender:    "male",
		Phone:     "+8562094938282",
		Address:   "ນາຄຳ",
	}
	users = append(users, item)

	id++
	item2 := User{
		ID:        id,
		FirstName: "sone",
		LastName:  "freestyle",
		Gender:    "male",
		Phone:     "+8562094938283",
		Address:   "ສາລາຄຳ",
	}
	users = append(users, item2)
}

func main() {
	e := echo.New()
	initUsers()

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		var req User

		if err := c.Bind(&req); err != nil {
			return err
		}

		id++
		req.ID = id
		users = append(users, req)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
