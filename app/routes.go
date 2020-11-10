package app

import (
	"git.hifx.in/crud_ops/handlers"
	"github.com/labstack/echo/v4"
)

//Routes creates the handlers, initializes the  muxes & sub-muxes & returns the final mux
func Routes(e *echo.Echo) {

	e.GET("/status", func(c echo.Context) error {
		c.Logger().Debug("debug me...\n")
		c.Logger().Info("info me...")
		c.Logger().Error("error me...")
		return c.JSON(200, "API is Running....")
	})

	//MYSQL CRUD
	e.GET("/users", handlers.ListUsers)        //To List all the users
	e.POST("/user", handlers.AddUser)          //To add a new user
	e.PUT("/user/:id", handlers.UpdateUser)    //To update record of existing user
	e.DELETE("/user/:id", handlers.DeleteUser) //To delete record of existing user
}
