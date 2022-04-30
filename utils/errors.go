package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// The variables aim to provide standardize error messages in the application
// to decouple error messages between DB tier and app tier
var (
	ErrorNotFound       = errors.New("not found")
	ErrorDuplicateEntry = errors.New("duplicate entry")
)

// Data structure for error reporting on HTTP calls
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// A function that return errors on HTTP calls
func Throws(c *gin.Context, status int, err error) {
	c.IndentedJSON(status, HTTPError{Code: status, Message: err.Error()})
}
