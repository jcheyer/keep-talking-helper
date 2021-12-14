package wires

import (
	"errors"
	"fmt"

	"github.com/jcheyer/keep-talking-helper/internals/deps"
)

type wires struct {
	wires  []string
	cmap   map[string]int
	bomb   deps.Bomb
	serial deps.Serial
}

func New(opts ...Option) (*wires, error) {
	w := &wires{
		wires: make([]string, 0),
		cmap:  make(map[string]int),
	}

	for _, opt := range opts {
		if err := opt(w); err != nil {
			return w, err
		}

	}

	return w, nil
}

func (w *wires) AddWire(s string) error {
	if !contains(validColors(), s) {
		return deps.ErrInvalidColor
	}
	w.cmap[s]++
	w.wires = append(w.wires, s)
	return nil
}

func (w *wires) Solve() (string, error) {
	switch len(w.wires) {
	case 3:
		if w.cmap["R"] == 0 {
			return fmt.Sprintf("Zweiten Draht durchtrennen (%s)", w.wires[1]), nil
		}
		if w.wires[2] == "W" {
			return fmt.Sprintf("Letzten Draht durchtrennen (%s)", w.wires[2]), nil
		}
		if w.cmap["B"] > 1 {
			for i := 2; i > 0; i-- {
				if w.wires[i] == "B" {
					return fmt.Sprintf("Letzten Blauen (%d) durchtrennen (%s)", i+1, w.wires[i]), nil
				}
			}
		}
		return fmt.Sprintf("Letzten Draht durchtrennen (%s)", w.wires[2]), nil
	case 4:
		if w.serial == nil {
			return "", deps.ErrSerNumUnknown
		}
		if w.cmap["R"] > 1 {
			if even, _ := w.serial.IsEven(); !even {
				for i := 3; i > 0; i-- {
					if w.wires[i] == "R" {
						return fmt.Sprintf("Letzten roten (%d) durchtrennen (%s)", i+1, w.wires[i]), nil
					}
				}
			}
		}
		if w.wires[3] == "G" && w.cmap["R"] == 0 {
			return fmt.Sprintf("Ersten Draht durchtrennen (%s)", w.wires[0]), nil
		}
		if w.cmap["B"] == 1 {
			return fmt.Sprintf("Ersten Draht durchtrennen (%s)", w.wires[0]), nil
		}
		if w.cmap["G"] > 1 {
			return fmt.Sprintf("Letzten Draht durchtrennen (%s)", w.wires[3]), nil
		}
		return fmt.Sprintf("Zweiten Draht durchtrennen (%s)", w.wires[1]), nil
	case 5:
		if w.serial == nil {
			return "", deps.ErrSerNumUnknown
		}
		if w.wires[4] == "S" {
			if even, _ := w.serial.IsEven(); !even {
				return fmt.Sprintf("Vierten Draht durchtrennen (%s)", w.wires[3]), nil
			}
		}
		if w.cmap["R"] == 1 && w.cmap["G"] > 1 {
			return fmt.Sprintf("Ersten Draht durchtrennen (%s)", w.wires[0]), nil
		}
		if w.cmap["S"] == 0 {
			return fmt.Sprintf("Zweiten Draht durchtrennen (%s)", w.wires[1]), nil
		}
		return fmt.Sprintf("Ersten Draht durchtrennen (%s)", w.wires[0]), nil

	case 6:
		if w.serial == nil {
			return "", deps.ErrSerNumUnknown
		}
		if w.cmap["G"] == 0 {
			if even, _ := w.serial.IsEven(); !even {
				return fmt.Sprintf("Dritten Draht durchtrennen (%s)", w.wires[2]), nil
			}
		}
		if w.cmap["G"] == 1 && w.cmap["W"] > 1 {
			return fmt.Sprintf("Vierten Draht durchtrennen (%s)", w.wires[3]), nil
		}
		if w.cmap["R"] == 0 {
			return fmt.Sprintf("Letzten Draht durchtrennen (%s)", w.wires[5]), nil
		}
		return fmt.Sprintf("Vierten Draht durchtrennen (%s)", w.wires[3]), nil

	}
	return "", errors.New("wrong number of wires")
}
