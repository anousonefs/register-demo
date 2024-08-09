package user

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// database
var users []*User

var id int

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		db,
	}
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func (s UserService) ListUsers(c echo.Context) error {

	query := "select id, first_name, last_name, gender, phone, address from users;"

	rows, err := s.db.QueryContext(c.Request().Context(), query)
	if err != nil {
		//todo return customer error
		return err
	}

	res := make([]User, 0)

	for rows.Next() {
		var i User
		rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Gender,
			&i.Phone,
			&i.Address,
		)

		res = append(res, i)
	}

	return c.JSON(200, res)
}

func (s UserService) CreateUser(c echo.Context) error {
	var req User
	if err := c.Bind(&req); err != nil {
		return err
	}

	insertSql := "insert into users(first_name, last_name, gender, phone, address) values($1, $2, $3, $4, $5);"

	result, err := s.db.ExecContext(
		c.Request().Context(),
		insertSql,
		req.FirstName,
		req.LastName,
		req.Gender,
		req.Phone,
		req.Address,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rowsAffected")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s UserService) UpdateUser(c echo.Context) (err error) {

	defer func() {
		if err != nil {
			logrus.Errorf("UpdateUser(): %v\n", err)
		}
	}()
	var req User
	if err = c.Bind(&req); err != nil {
		// customer error
		return err
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		// customer error
		return err
	}

	updateSql := `UPDATE users
								SET 
										first_name = $1,
										last_name = $2,
										gender = $3,
										phone = $4,
										address = $5
								WHERE 
										id = $6;
	`

	result, err := s.db.ExecContext(
		c.Request().Context(),
		updateSql,
		req.FirstName,
		req.LastName,
		req.Gender,
		req.Phone,
		req.Address,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rowsAffected")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (s UserService) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// customer error
		return err
	}

	deleteSql := "delete from users where id = $1"

	result, err := s.db.ExecContext(c.Request().Context(), deleteSql, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rowsAffected")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
