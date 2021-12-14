package complicatedwires

import "fmt"

type wire struct {
	color string
	led   bool
	star  bool
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (w *wire) String() string {
	return fmt.Sprintf("%d%s%d", bool2int(w.led), w.color, bool2int(w.star))
}

func (w *wire) SimplifiedString() string {
	color := w.color
	if w.color == "RW" {
		color = "R"
	}
	if w.color == "BW" {
		color = "B"
	}
	return fmt.Sprintf("%d%s%d", bool2int(w.led), color, bool2int(w.star))
}

// https://github.com/cnguyen-uk/Keep-Talking-and-Nobody-Explodes/blob/master/src/complicated_wires.py

func (w *wire) Solve() string {
	m := map[string]string{
		"0W0": "D",
		"0W1": "D",
		"1W0": "N",
		"1W1": "B",

		"0B0": "S",
		"0B1": "N",
		"1B0": "P",
		"1B1": "P",

		"0R0": "S",
		"0R1": "D",
		"1R0": "B",
		"1R1": "B",

		"0BR0": "S",
		"0BR1": "N",
		"1BR0": "S",
		"1BR1": "N",
	}

	return m[w.SimplifiedString()]
}
