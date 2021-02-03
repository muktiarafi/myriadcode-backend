package controllers

import (
	"errors"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) error {

	return errors.New(http.StatusText(http.StatusBadRequest))
}
