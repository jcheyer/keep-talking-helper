package serial

import (
	"strconv"

	"github.com/jcheyer/keep-talking-helper/internals/deps"
)

type serial struct {
	serial  string
	isEven  bool
	isValid bool
}

type Option func(s *serial) error

func New(opts ...Option) (*serial, error) {
	s := &serial{}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return s, err
		}

	}

	return s, nil
}

func (ser *serial) SetSerial(s string) error {
	if len(s) == 0 {
		return deps.ErrSerNumInvalid
	}

	cn := s[len(s)-1]
	n, err := strconv.Atoi(string(cn))
	if err != nil {
		return deps.ErrSerNumInvalid
	}
	ser.serial = s
	ser.isEven = n%2 == 0
	ser.isValid = true
	return nil
}

func WithSerial(s string) Option {
	return func(ser *serial) error {
		return ser.SetSerial(s)
	}
}

func (ser *serial) IsEven() (bool, error) {
	if !ser.isValid {
		return false, deps.ErrSerNumInvalid
	}
	return ser.isEven, nil
}
func (ser *serial) IsValid() bool {
	return ser.isValid
}
