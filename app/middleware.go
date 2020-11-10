package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"git.hifx.in/crud_ops/domain"

	// query "git.hifx.in/lens/querybuilder/redshift"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Http Errors from Middleware
var (
	ErrForbiden         = domain.NewHTTPError(http.StatusForbidden, "You are not authorized to perform this action")
	ErrInvalidAdminID   = "Invalid admin ID"
	ErrInvalidAdminName = "Invalid admin Name"
)

// CORSAllower sets required headers for enabling cross origin header
func CORSAllower() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderXRequestedWith,
			echo.HeaderAuthorization,
			"X-Lens-Debug-Vars",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		ExposeHeaders: []string{echo.HeaderXRequestID},
		MaxAge:        600,
	})

}

var prefix string
var reqid uint64

// SetDebugVars sets debug variables in context
func SetDebugVars(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get(domain.SUPERADMIN) != nil {
			debugVars := strings.Split(c.Request().Header.Get("X-Lens-Debug-Vars"), ",")
			for _, v := range debugVars {
				c.Set(string(domain.DebugVar(v)), true)
			}
		}
		err := next(c)
		// Here, we set the response headers for the debugging variables
		if c.Get(string(domain.ADDQUERY)) != nil && c.Get(domain.QUERY) != nil {
			c.Response().Header().Add(domain.QUERY, c.Get(domain.QUERY).(string))
		}

		return err
	}
}

// OpenLogFile opens the log fileName to track the log data
func OpenLogFile(fileName string) (io.Writer, error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return nil, err
	}

	return io.MultiWriter(f), nil
}

func compare(match string, matchWith []string) bool {
	for _, m := range matchWith {
		if m == match {
			return true
		}
	}

	return false
}

// ErrorHandler handles the customized error handler
func ErrorHandler(err error, c echo.Context) {
	// *echo.HTTPError is useful if you don't have any internel error but still wants
	// send an http response as an error.
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, domain.Error{
			Message: he.Error(),
		})
	}

	// If you like to handle a fully customizable error
	if derr, ok := err.(*domain.Error); ok {
		c.JSON(derr.Code, derr)
	}

}

// Recover returns a Recover middleware with config.
func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				r := recover()
				if r != nil {
					r = fmt.Errorf("%v", r)
					c.Logger().Error(r)
					c.JSON(http.StatusInternalServerError, domain.Error{
						Message: "Sorry, Something went wrong. ",
					})
				}
			}()

			return next(c)
		}
	}
}
