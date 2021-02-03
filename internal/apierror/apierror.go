package apierror

import (
	"fmt"
	"net/http"
)

type Error struct {
	StatusCode int
	Err        error
	Kind       error
	Warning    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("API Error: %s", e.Err.Error())
}

func NewInternalServerError(err error, warning string) *Error {
	return &Error{http.StatusInternalServerError, err, InternalServerError, warning}
}
