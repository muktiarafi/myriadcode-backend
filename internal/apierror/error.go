package apierror

import "net/http"

type ErrString string

func (e ErrString) Error() string {
	return string(e)
}

var (
	BadRequestError     = ErrString(http.StatusText(http.StatusBadRequest))
	InternalServerError = ErrString(http.StatusText(http.StatusInternalServerError))
	UnathorizedError    = ErrString(http.StatusText(http.StatusUnauthorized))
	NotFoundError       = ErrString(http.StatusText(http.StatusNotFound))
)
