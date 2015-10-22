package utilities

import (
	"errors"
	"log"
	"net/http"
	"runtime"
)

type Error struct {
	InnerError error
	Buf        []byte
	Status     int
	Annotation string
}

func NewAppError(optionalMessage string, optionalInnerError error, captureStackTrace bool, contextDependantStatusInt int) *Error {
	err := Error{
		InnerError: optionalInnerError,
		Buf:        nil,
		Status:     contextDependantStatusInt,
		Annotation: optionalMessage,
	}
	if captureStackTrace {
		err.Buf = make([]byte, 1<<16)
		runtime.Stack(err.Buf, false)
	}

	return &err
}

func ServerError(err error) *Error {
	if err == nil {
		return nil
	}

	apperr := NewAppError("", err, true, 500)
	return apperr
}

func ClientError(message string, status int) *Error {
	apperr := NewAppError("", errors.New(message), false, status)
	return apperr
}

type Handler func(w http.ResponseWriter, req *http.Request) *Error

func (handler Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	apperr := handler(w, req)
	if apperr != nil {
		var message string
		if apperr.Status >= 500 {
			message = "Internal Server Error"
		} else {
			message = apperr.Error()
		}
		log.Printf("ERROR (%d): %v\n%s", apperr.Status, apperr.InnerError, apperr.Buf)
		http.Error(w, message, apperr.Status)
	}
}

func (apperr *Error) Error() string {
	if apperr.Annotation == "" {
		if apperr.InnerError != nil {
			return apperr.InnerError.Error()
		} else {
			return "Unspecified error"
		}
	} else {
		if apperr.InnerError != nil {
			return apperr.Annotation + " :\n" + apperr.InnerError.Error()
		} else {
			return apperr.InnerError.Error()
		}
	}
}
