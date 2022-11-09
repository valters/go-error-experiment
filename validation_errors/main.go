package validation_errors

import "errors"

var SqlNoSuchRow = errors.New("row with id = 404 not found")
var SqlConnectionFailed = errors.New("could not connect to db localhost:4242")

var MemberNotFound = &MemberNotFoundType{"", nil}
var MemberGenericError = &MemberGenericErrorType{"", nil}

type MemberNotFoundType wrapError
type MemberGenericErrorType wrapError

type wrapError struct {
	msg string
	err error
}

// Is was called on value of MemberNotFoundType and the caller wants to know if this is a MemberNotFound ? Why yes, yes it is.
func (e *MemberNotFoundType) Is(y error) bool {
	if y == nil {
		return false
	}
	if y == MemberNotFound {
		return true
	}
	return false
}

// Is was called on value of MemberGenericErrorType and the caller wants to know if this is a MemberGenericError ? Why yes, yes it is.
func (e *MemberGenericErrorType) Is(y error) bool {
	if y == nil {
		return false
	}
	if y == MemberGenericError {
		return true
	}
	return false
}

func (e *wrapError) Error() string {
	return e.msg + " -> " + safeError(e.err)
}

func safeError(e error) string {
	if e == nil {
		return ""
	}
	return e.Error()
}

func (e *MemberNotFoundType) Error() string {
	return e.msg + " -> " + safeError(e.err)
}

func (e *MemberGenericErrorType) Error() string {
	return e.msg + " -> " + safeError(e.err)
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func (e *MemberNotFoundType) Unwrap() error {
	return e.err
}

func (e *MemberGenericErrorType) Unwrap() error {
	return e.err
}

func MkMemberNotFound(msg string, err error) *MemberNotFoundType {
	return &MemberNotFoundType{msg, err}
}
