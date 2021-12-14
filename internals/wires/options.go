package wires

import "github.com/jcheyer/keep-talking-helper/internals/deps"

type Option func(w *wires) error

func WithBomb(b deps.Bomb) Option {
	return func(w *wires) error {
		w.bomb = b
		return nil
	}
}

func WithSerial(ser deps.Serial) Option {
	return func(w *wires) error {
		w.serial = ser
		return nil
	}
}
