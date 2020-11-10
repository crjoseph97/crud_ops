package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Global Error
const (
	ErrInternalServer         = "Sorry, something went wrong."
	ErrInvalidAdmin           = "Admin not found"
	ErrForbidden              = "You are not authorized to perform this action"
	ErrInvalidDashboard       = "Dashboard not found"
	ErrInvalidChart           = "Chart not found"
	ErrInvalidCredentials     = "Invalid email / password"
	ErrEmptyPassword          = "Password can not be empty"
	ErrEmailAlreadyInUse      = "Email already in use"
	ErrInvalidRequest         = "Invalid request"
	ErrInvalidRoleID          = "Role not found"
	ErrInvalidCode            = "Invalid code"
	ErrAuthorization          = "You are not authorized to perform this action"
	ErrInvalidOldPassword     = "Incorrect old password"
	ErrDimensionNotFound      = "Dimension not found"
	ErrAdminsNotFound         = "No admins found which satisfies the given filters"
	ErrPhoneNotFound          = "No phone number for this user"
	ErrAccountIDNotFound      = "Invalid account id for this user"
	ErrLoginExpired           = "Login expired, please login again"
	ErrDashboardAuthorization = "You do not have authorization on this dashboard"
	ErrTemplateNotFound       = "Template not found for query"
	ErrFunnelNotFound         = "Funnel not found"
	ErrCampaignNotFound       = "Campaign not found"
	ErrCustomQueryNotFound    = "Custom Query not found"
	ErrInvalidQuery           = "The query provided is invalid"
	ErrCustomEvent            = "The custom event you are trying to delete is used for charts in dashboard"
	ErrInvalidUser            = "Profile does not exist"
	ErrValidation             = "you request has been failed in our validation process. please resolve the given errors."
)

// Account Error
const (
	ErrAccountsNotFound = "No account found"
)

// Error defines the error response model
type Error struct {
	Code    int         `json:"-"`
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Error marshals the error object to JSON
func (e *Error) Error() string {
	s, err := json.Marshal(*e)
	if err != nil {
		return fmt.Sprintf("Marshall Error: %s in Error: %s", err.Error(), e.Error())
	}
	return RemoveUnwantedSpaces(string(s))
}

// Unwrap grabs the error object alone
func (e *Error) Unwrap() error {
	if err, ok := e.Errors.(error); ok {
		return err
	}
	return nil
}

// NewHTTPError ...
func NewHTTPError(code int, errMessage ...interface{}) *Error {
	err := &Error{
		Code:    code,
		Message: http.StatusText(code),
	}

	if errMessage != nil && len(errMessage) > 0 {
		if errMessage[0] != nil {
			err.Message = errMessage[0]
		}

		if len(errMessage) > 1 && errMessage[1] != nil {
			if e, ok := errMessage[1].(error); ok {
				err.Errors = e
			}
		}
	}

	return err
}

// ValidationErr ..
type ValidationErr map[string]string

func (verr ValidationErr) Error() string {
	s, _ := json.MarshalIndent(verr, "", "  ")
	return RemoveUnwantedSpaces(string(s))
}
