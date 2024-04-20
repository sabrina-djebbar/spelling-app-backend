package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// Error is a formatted error response.
type Error struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`

	InnerError error `json:"inner_error,omitempty"`
}

type hasInnerError struct {
	InnerError *Error `json:"inner_error,omitempty"`
}

type jsonAlias Error

// ErrChain allows for multiple separate operations to run. If any fail, check
// will return an error containing the results.
func ErrChain() (chain func(...error), check func() error) {
	var (
		mux    sync.Mutex
		stored error
	)

	return func(errs ...error) {
			for _, err := range errs {
				if err != nil {
					mux.Lock()
					stored = errors.Join(stored, err)
					mux.Unlock()
				}
			}
		},
		func() error { return stored }
}

func (e *Error) UnmarshalJSON(data []byte) error {
	ja := &jsonAlias{}
	_ = json.Unmarshal(data, ja)

	hie := &hasInnerError{}
	err := json.Unmarshal(data, hie)
	if err != nil {
		return err
	}

	*e = Error{
		Code:    ja.Code,
		Message: ja.Message,
		Details: ja.Details,
	}

	if hie.InnerError != nil {
		e.InnerError = *hie.InnerError
	}

	return nil
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonAlias{
		Code:    e.Code,
		Message: e.Message,
		Details: e.Details,

		InnerError: e.InnerError,
	})
}

// Details is a short code for a map
type Details map[string]interface{}

type WrapOption func(Error) Error

func WithCode(code string) WrapOption {
	return func(e Error) Error {
		e.Code = code
		return e
	}
}

func WithDetails(details Details) WrapOption {
	return func(e Error) Error {
		// merge details maps if possible, else replace.
		switch oldDetails := e.Details.(type) {
		case Details:
			newDetails := make(Details)

			for k, v := range oldDetails {
				newDetails[k] = v
			}
			for k, v := range details {
				newDetails[k] = v
			}

			e.Details = newDetails

		default:
			e.Details = details
		}

		return e
	}
}

func WithMessage(message string) WrapOption {
	return func(e Error) Error {
		e.Message = message
		return e
	}
}

// WithMessagef applies [fmt.Printf]-like formatting to the error message. It is
// not recommended to include IDs in the message to make structured log searches
// more effective - use werr.WithDetails instead.
func WithMessagef(format string, a ...interface{}) WrapOption {
	return func(e Error) Error {
		e.Message = fmt.Sprintf(format, a...)
		return e
	}
}

func New(message string, opts ...WrapOption) Error {
	err := Error{Message: message}

	for _, opt := range opts {
		err = opt(err)
	}
	return err
}

func Wrap(err error, opts ...WrapOption) Error {
	errResponse := Error{InnerError: err, Message: err.Error()}
	details := Details{}

	// extract any inner werr details
	if innerWErr := (Error{}); errors.As(err, &innerWErr) {
		errResponse.Code = innerWErr.Code

		if innerWErr.Details != nil {
			details["inner_details"] = innerWErr.Details
		}
	}

	if len(details) != 0 {
		errResponse.Details = details
	}

	for _, opt := range opts {
		errResponse = opt(errResponse)
	}
	return errResponse
}

func (e Error) Error() string {
	msg := e.Message
	if e.InnerError != nil {
		embedMsg := e.InnerError.Error()
		if embedMsg != msg {
			msg += ": " + embedMsg
		}
	}
	return msg
}

func (e Error) Unwrap() error {
	return e.InnerError
}
