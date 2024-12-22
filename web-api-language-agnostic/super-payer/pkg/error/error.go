package error

import "fmt"

type Kind string

const (
	ErrKindNotFound        = Kind("not_found")
	ErrKindInvalidArgument = Kind("invalid_argument")
	ErrKindInternal        = Kind("internal")
)

type PkgError interface {
	GetKind() Kind
	Error() string
}

type Error struct {
	kind Kind
	msg  string
	err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"kind: %s, msg: %s, err: %s",
		e.kind,
		e.msg,
		e.err.Error())
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) GetKind() Kind {
	return e.kind
}

func NewPkgErrorInternal(msg string, err error) Error {
	return Error{
		kind: ErrKindInternal,
		msg:  msg,
		err:  err,
	}
}

func NewPkgErrorInvalidArgument(msg string, err error) Error {
	return Error{
		kind: ErrKindInvalidArgument,
		msg:  msg,
		err:  err,
	}
}

func NewPkgErrorNotFound(msg string, err error) Error {
	return Error{
		kind: ErrKindNotFound,
		msg:  msg,
		err:  err,
	}
}
