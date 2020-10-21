package errno

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

func ErrorJSON(err error) *gin.H {
	if err == nil {
		return &gin.H{"code": OK.Code, "message": OK.Message}
	}

	switch typed := err.(type) {
	case *Err:
		return &gin.H{"code": typed.Code, "message": typed.Message}
	case *Errno:
		return &gin.H{"code": typed.Code, "message": typed.Message}
	default:
	}

	return &gin.H{"code": InternalServerError.Code, "message": err.Error()}
}
