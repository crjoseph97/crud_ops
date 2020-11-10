/*
Package utils provides a handfull of utility tools  which provides more flexiblity for an application.
*/
package utils

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

var log echo.Logger

// Init initializes the Utils package with needed dependencies
func Init(logger echo.Logger) {
	log = logger
}

// Close an io.Closer object by logging it's error.
//  Example:
// 		defer Close(closer,SrcMessage)
// 		defer Close(closer)
// 		defer Close(closer,"dimension.reader")
func Close(ctx echo.Context, closer io.Closer, SrcMessage ...interface{}) {
	if err := closer.Close(); err != nil {
		if ctx != nil {
			log.Error(ctx, err, fmt.Sprintf("closer(%#v)\n", closer), SrcMessage)
		}

		log.Error(err, fmt.Sprintf("closer(%#v)", closer), SrcMessage)
	}
}

type rollBack interface {
	Rollback() error
}

// RollBack a transactions by logging if any error
func RollBack(ctx echo.Context, tx rollBack, SrcMessage ...string) {
	if err := tx.Rollback(); err != nil {
		if ctx != nil {
			log.Error(ctx, err, fmt.Sprintf("closer(%#v)\n", tx), SrcMessage)
		}

		log.Error(err, fmt.Sprintf("closer(%#v)", tx), SrcMessage)
	}
}
