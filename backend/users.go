package backend

import (
	"net/http"

	"git.hifx.in/crud_ops/domain"
	"git.hifx.in/crud_ops/utils"

	"github.com/labstack/echo"
)

//UsersRepo ...
type UsersRepo interface {
	ListUsers(c echo.Context, query string) ([]domain.User, error)
	AddUser(c echo.Context, query string) error
	UpdateUser(c echo.Context, query string) error
	DeleteUser(c echo.Context, query string) error
}

type users struct {
	*Infra
}

// ListUsers function returns the details of uses
func (u users) ListUsers(c echo.Context, query string) ([]domain.User, error) {
	list := make([]domain.User, 0)
	rows, err := u.DBr.Query(query)
	if err != nil {
		c.Logger().Error(err)
		return list, domain.NewHTTPError(http.StatusInternalServerError, ErrQueryExecutionFailed)
	}

	defer utils.Close(c, rows, ErrSQLConnection)

	for rows.Next() {
		var temp domain.User
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Mobile, &temp.Address, &temp.Department)
		if err != nil {
			c.Logger().Error(err)
			return list, domain.NewHTTPError(http.StatusInternalServerError, ErrScanFailed)
		}

		list = append(list, temp)
	}
	return list, err
}

// AddUser function adds new records
func (u users) AddUser(c echo.Context, query string) error {
	_, err := u.DBw.Exec(query)
	if err != nil {
		c.Logger().Error(err)
		return domain.NewHTTPError(http.StatusInternalServerError, ErrQueryExecutionFailed)
	}
	return nil
}

// UpdateUser function updates the record
func (u users) UpdateUser(c echo.Context, query string) error {
	_, err := u.DBw.Exec(query)
	if err != nil {
		c.Logger().Error(err)
		return domain.NewHTTPError(http.StatusInternalServerError, ErrQueryExecutionFailed)
	}
	return nil
}

// DeleteUser function deletes the record
func (u users) DeleteUser(c echo.Context, query string) error {
	_, err := u.DBw.Exec(query)
	if err != nil {
		c.Logger().Error(err)
		return domain.NewHTTPError(http.StatusInternalServerError, ErrQueryExecutionFailed)
	}
	return nil
}
