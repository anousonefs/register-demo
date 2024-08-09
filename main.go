package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// database
var users []*User

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
	users = append(users, &item)

	id++
	item2 := User{
		ID:        id,
		FirstName: "sone",
		LastName:  "freestyle",
		Gender:    "male",
		Phone:     "+8562094938283",
		Address:   "ສາລາຄຳ",
	}
	users = append(users, &item2)
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
		users = append(users, &req)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		var req User
		if err := c.Bind(&req); err != nil {
			// customer error
			return err
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// customer error
			return err
		}

		for _, i := range users {
			if i.ID == id {
				i.FirstName = req.FirstName
				i.LastName = req.LastName
				i.Gender = req.Gender
				i.Phone = req.Phone
				i.Address = req.Address
			}
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			// customer error
			return err
		}

		//id: 4
		// 1, 2, 3, 4, 5, 6, 7
		// users[:id] = 1, 2, 3
		//users[id+1:] = 5, 6, 7

		isFound := false
		for index, i := range users {
			if i.ID == id {
				isFound = true
				users = append(users[:index], users[index+1:]...)
			}
		}

		if !isFound {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "not found",
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})

	})

	e.Logger.Fatal(e.Start(":8080"))
}
