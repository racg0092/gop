package gop

import (
	"errors"
)

var (
	ErrUnabelToAuthenticate = errors.New("unable to authenticate password or username does not match")
	ErrShortPassword        = errors.New("password is less than 16 characters long. it needs to be longer")
	ErrPwnedPassword        = GopError{message: "this password has been exposed previously on a data breach"}
)

type GopError struct {
	appended string
	message  string
}

func (e GopError) Error() string {
	msg := e.message + " " + e.appended
	e.appended = ""
	return msg
}

func (e *GopError) append(s string) {
	e.appended = e.appended + s
}
