package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"git.hifx.in/crud_ops/domain"
	"github.com/labstack/echo/v4"
)

// Error constants for handlers
const (
	ErrInvalidFilters = "Invalid value in `filters`"
	ErrMissingUserID  = "User ID path parameter missing"
)

type response struct {
	Data    interface{} `json:"data,omitempty"`
	Query   interface{} `json:"query,omitempty"`
	ESIndex interface{} `json:"es_index,omitempty"`
}

// StatusOk ...
func StatusOk(c echo.Context, resp interface{}) error {
	r := response{Data: resp}

	if c.Get(string(domain.ADDQUERY)) != nil {
		r.Query = c.Get(domain.QUERY)
	}

	return c.JSON(http.StatusOK, r)
}

// ReplacePositionalParamsInQuery replaces the positional parameters in the query
// with the corresponding values for logging purpose
// example:
//	query = select * from table where id = ? and category = ?
//	params = [1, "c"]
//	returns select * from table where id = 1 and category = 'c'
func ReplacePositionalParamsInQuery(query string, params ...interface{}) string {
	// TODO: This is a naive method for replacement. A better implementation can be added as required. (-_-) zzz
	for _, param := range params {
		query = strings.Replace(query, "?", getValueAsString(param), 1)
	}
	return query
}

var spaceRemover = regexp.MustCompile("\\s\\s+")

// RemoveUnwantedSpaces replaces multiple adjacent space with single
func RemoveUnwantedSpaces(str string) string {
	return spaceRemover.ReplaceAllString(str, " ")
}

func getValueAsString(value interface{}) string {
	if value == nil {
		return "NULL"
	}
	switch typedValue := value.(type) {
	case
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", typedValue)
	case float32, float64:
		return fmt.Sprintf("%f", typedValue)
	case *string:
		return fmt.Sprintf("'%s'", *typedValue)
	default:
		return fmt.Sprintf("'%v'", typedValue)
	}
}
