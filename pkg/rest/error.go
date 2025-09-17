package rest

import (
	"encoding/json"
	"errors"
)

var (
	ErrConvID               = errors.New("converting id error")
	ErrParseQuery           = errors.New("parsing query params error")
	ErrRequiredValue        = errors.New("missing required value")
	ErrInternalServer       = errors.New("internal server error")
	ErrMarshal              = errors.New("cannot marhsal")
	ErrUnmarshal            = errors.New("cannot unmarshal")
	ErrRequestBodyReading   = errors.New("cannot read request body")
	ErrNotFound             = errors.New("not found")
	ErrUnsupportedMediaType = errors.New("only application/json is allowed")
	ErrBadRequest           = errors.New("invalid json message received")
	ErrEmptyBody            = errors.New("request must contain elements in array")
	ErrConstraintViolation  = errors.New("constraint violation")
	ErrMissingQuery         = errors.New("missing query params")
	ErrEmptyValue           = errors.New("value from query cannot be empty")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func ExtractRestErrorFromBody(body []byte) (*Error, error) {
	var restErr *Error
	err := json.Unmarshal(body, &restErr)
	if err != nil {
		return nil, err
	}
	if restErr == nil || restErr.Code == 0 {
		return nil, errors.New(string(body))
	}

	return restErr, nil
}
