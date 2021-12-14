package complicatedwires

import (
	"fmt"

	"github.com/jcheyer/keep-talking-helper/internals/deps"
	"github.com/jcheyer/keep-talking-helper/internals/serial"
)

type cw struct {
	wires             []wire
	serial            deps.Serial
	parallelPortKnown bool
	parallelPort      bool
	batteryCount      int
	bomb              deps.Bomb
}

func New(opts ...Option) *cw {
	cw := &cw{
		wires:        make([]wire, 0),
		batteryCount: -1,
	}

	for _, opt := range opts {
		opt(cw)
	}

	if cw.bomb != nil {
		b := cw.bomb

		sn, err := b.GetSerial()
		if err != nil && sn != "" {
			_ = cw.SetSerial(sn)
		}
	}

	return cw
}

func (cw *cw) Run() {

}

func (cw *cw) SetSerial(s string) error {
	ser, err := serial.New(serial.WithSerial(s))
	if err != nil {
		return err
	}

	cw.serial = ser
	return nil
}

func (cw *cw) SetBatteryCount(c int) {
	cw.batteryCount = c
}

func (cw *cw) SetHasParallelPort() {
	cw.parallelPort = true
	cw.parallelPortKnown = true
}

func (cw *cw) AddWire(color string, led bool, star bool) error {
	color, err := normalizeColor(color)
	if err != nil {
		return err
	}

	cw.wires = append(cw.wires, wire{color, led, star})
	return nil
}

func (cw *cw) SolveWire(w wire) (bool, error) {
	action := w.Solve()
	switch action {
	case "D":
		return true, nil
	case "N":
		return false, nil
	case "S":
		if cw.serial == nil || !cw.serial.IsValid() {
			return false, deps.ErrSerNumUnknown
		}
		return cw.serial.IsEven()
	case "P":
		if !cw.parallelPortKnown {
			return false, deps.ErrParallelPortUnknown
		}
		return cw.parallelPort, nil

	case "B":
		if cw.batteryCount == -1 {
			return false, deps.ErrBatteriesUnknown
		}
		if cw.batteryCount >= 2 {
			return true, nil
		}
		return false, nil
	}

	return false, fmt.Errorf("unknown action: %s", action)
}
