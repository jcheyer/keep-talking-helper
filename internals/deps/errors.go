package deps

import "errors"

var (
	ErrBatteriesUnknown    = errors.New("battery count unknown")
	ErrSerNumUnknown       = errors.New("serialnumber unknown")
	ErrSerNumInvalid       = errors.New("serialnumber invalid")
	ErrParallelPortUnknown = errors.New("parallelport unknown")
	ErrInvalidColor        = errors.New("invalid color")
)
