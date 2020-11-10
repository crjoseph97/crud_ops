/*
Package app provides Infrastructure setup for the application. Infrastrucre contains
   External session objects for databases,aws
   Exposing Restful API's
   Middlewares to handle API's
   Initializing Configuration & Backend services.
*/
package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"git.hifx.in/crud_ops/backend"
	"git.hifx.in/crud_ops/conf"
	"git.hifx.in/crud_ops/utils"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func init() {
	// Initialize application
	err := Init()
	if err != nil {
		log.Println("Error: initializing app:- ", err.Error())
		os.Exit(0)
	}
}

/*
Init initializes the application by setting up configurations like
database connections.
*/
func Init() error {
	Dbr, Dbw, err := connectSQL()
	if err != nil {
		return fmt.Errorf("Failed to connect MySQL\nError: %s", err.Error())
	}
	infra := &backend.Infra{
		DBr: Dbr,
		DBw: Dbw,
	}
	backend.Init(infra)
	return nil
}

func connectSQL() (Dbr, Dbw *sqlx.DB, err error) {
	Dbr, err = backend.ConnectDB(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.Current.UserName,
			conf.Current.Password,
			conf.Current.Host,
			conf.Current.DBPort,
			conf.Current.Database,
		),
		conf.Current.MaxActive,
		conf.Current.MaxIdle,
	)

	if err != nil {
		return nil, nil, err
	}

	Dbw, err = backend.ConnectDB(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.Current.UserName,
			conf.Current.Password,
			conf.Current.Host,
			conf.Current.DBPort,
			conf.Current.Database,
		),
		conf.Current.MaxActive,
		conf.Current.MaxIdle,
	)
	if err != nil {
		return nil, nil, err
	}

	return Dbr, Dbw, err
}

// Run runs the entire application
func Run() {
	e := echo.New()
	if conf.Current.Mode != conf.MODEDEV {
		f, err := os.OpenFile(conf.Current.Combined, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Test fail 1", conf.Current.Combined)
			return
		}
		defer f.Close()
		e.Logger.SetOutput(f)
		e.HideBanner = true
	}
	e.Debug = true

	// Log init for utils package
	utils.Init(e.Logger)

	e.HTTPErrorHandler = ErrorHandler
	// Inits Routes

	Routes(e)
	s := &http.Server{
		Addr:         conf.Current.Port,
		WriteTimeout: time.Duration(conf.Current.WriteTimeout) * time.Second,
	}
	if conf.Current.Mode != conf.MODEDEV {
		fmt.Println("----------------------------")
		fmt.Println("Listening to port ", conf.Current.Port)
		fmt.Println("----------------------------")
	}
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	gracefulHandle(e)
}

func gracefulHandle(e *echo.Echo) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Current.GraceTimeout)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
