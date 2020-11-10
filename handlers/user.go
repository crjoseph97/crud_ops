package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.hifx.in/crud_ops/backend"
	"git.hifx.in/crud_ops/domain"

	"github.com/labstack/echo"
	// query "git.hifx.in/lens/querybuilder/redshift"
)

// User holds for the user handlers method from routes
type User struct{}

// ListUsers retrieves user list
func ListUsers(c echo.Context) error {
	qb := domain.UserBuilder{}
	err := json.Unmarshal([]byte(c.FormValue("dep")), &qb.Department)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	query, err := domain.UserQueryBuilder(qb)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	fmt.Println(query)
	data, err := backend.Users.ListUsers(c, query)
	if err != nil {
		return err
	}

	response := map[string]interface{}{}
	response["values"] = data
	var res = "This is response data"
	response["value"] = res
	return StatusOk(c, response)
}

// AddUser adds a new user
func AddUser(c echo.Context) error {
	qb := new(domain.UserBuilder)
	if err := c.Bind(qb); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	query, err := domain.CreateQueryBuilder(*qb)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	err = backend.Users.AddUser(c, query)
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, "User added successfuly")
}

// UpdateUser updates details of a user
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	qb := new(domain.UserBuilder)
	if err := c.Bind(qb); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	query, err := domain.UpdateQueryBuilder(*qb, id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	fmt.Println(query)
	err = backend.Users.UpdateUser(c, query)
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, "User values updated")
}

// DeleteUser deletes details of a user
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	query, err := domain.DeleteQueryBuilder(id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	fmt.Println(query)
	err = backend.Users.DeleteUser(c, query)
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, "User record deleted")
}
