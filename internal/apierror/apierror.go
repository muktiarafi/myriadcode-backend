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

func checkWarning(warning *string, value string) {
	if len(*warning) == 0 {
		*warning = value
	}
}

func NewInternalServerError(err error, warning string) *Error {
	checkWarning(&warning, InternalServerError.Error())
	return &Error{http.StatusInternalServerError, err, InternalServerError, warning}
}

func NewBadRequestError(err error, warning string) *Error {
	kind := BadRequestError
	checkWarning(&warning, kind.Error())
	return &Error{http.StatusBadRequest, err, kind, warning}
}
