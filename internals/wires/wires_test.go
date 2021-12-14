package wires

import (
	"testing"

	"github.com/jcheyer/keep-talking-helper/internals/serial"
	"github.com/stretchr/testify/assert"
)

func Test3(t *testing.T) {
	tests := []struct {
		name       string
		w1, w2, w3 string
		expect     string
	}{
		{
			name: "kein rot -> zweiter Draht",
			w1:   "B", w2: "B", w3: "B",
			expect: "Zweiten Draht durchtrennen (B)",
		},
		{
			name: "letzer weiß -> letzter Draht",
			w1:   "R", w2: "B", w3: "W",
			expect: "Letzten Draht durchtrennen (W)",
		},
		{
			name: "letzer weiß -> letzter Draht",
			w1:   "B", w2: "R", w3: "B",
			expect: "Letzten Blauen (3) durchtrennen (B)",
		},
		{
			name: "Andernfalls -> letzer",
			w1:   "B", w2: "W", w3: "R",
			expect: "Letzten Draht durchtrennen (R)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := New()
			assert.NoError(t, err)
			assert.NoError(t, w.AddWire(tt.w1))
			assert.NoError(t, w.AddWire(tt.w2))
			assert.NoError(t, w.AddWire(tt.w3))
			s, err := w.Solve()
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, s)
		})
	}
}

func Test4(t *testing.T) {
	tests := []struct {
		name           string
		w1, w2, w3, w4 string
		serial         string
		expect         string
	}{
		{
			name: "mehr als ein rot + ungerade serial -> letzter roter",
			w1:   "B", w2: "B", w3: "R", w4: "R",
			serial: "gfsgfds1",
			expect: "Letzten roten (4) durchtrennen (R)",
		},
		{
			name: "letzter Gelb, kein rot -> erster draht",
			w1:   "W", w2: "B", w3: "W", w4: "G",
			expect: "Ersten Draht durchtrennen (W)",
		},
		{
			name: "einer blau -> erster draht",
			w1:   "R", w2: "B", w3: "R", w4: "W",
			serial: "gfsgfds2", // siehe oben
			expect: "Ersten Draht durchtrennen (R)",
		},
		{
			name: "mehr als ein gelb -> letzter draht",
			w1:   "G", w2: "G", w3: "R", w4: "W",
			expect: "Letzten Draht durchtrennen (W)",
		},
		{
			name: "ansonsten zweiter draht",
			w1:   "G", w2: "R", w3: "R", w4: "W",
			expect: "Zweiten Draht durchtrennen (R)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serNum := tt.serial
			if serNum == "" {
				serNum = "2"
			}
			ser, err := serial.New(serial.WithSerial(serNum))
			assert.NoError(t, err)
			w, err := New(WithSerial(ser))

			assert.NoError(t, err)
			assert.NoError(t, w.AddWire(tt.w1))
			assert.NoError(t, w.AddWire(tt.w2))
			assert.NoError(t, w.AddWire(tt.w3))
			assert.NoError(t, w.AddWire(tt.w4))
			s, err := w.Solve()
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, s)
		})
	}
}

func Test5(t *testing.T) {
	tests := []struct {
		name               string
		w1, w2, w3, w4, w5 string
		serial             string
		expect             string
	}{
		{
			name: "letzter schwarz + ungerade serial -> vierter draht",
			w1:   "B", w2: "B", w3: "R", w4: "R", w5: "S",
			serial: "gfsgfds1",
			expect: "Vierten Draht durchtrennen (R)",
		},
		{
			name: "einer rot und mehr als zwei gelbe -> erster draht",
			w1:   "R", w2: "G", w3: "G", w4: "G", w5: "S",
			serial: "gfsgfds2",
			expect: "Ersten Draht durchtrennen (R)",
		},
		{
			name: "kein schwarzer -> zweiten draht",
			w1:   "R", w2: "B", w3: "R", w4: "W", w5: "W",
			serial: "gfsgfds2",
			expect: "Zweiten Draht durchtrennen (B)",
		},
		{
			name: "ansonsten -> erster draht",
			w1:   "G", w2: "G", w3: "W", w4: "W", w5: "S",
			serial: "gfsgfds2",
			expect: "Ersten Draht durchtrennen (G)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serNum := tt.serial
			if serNum == "" {
				serNum = "2"
			}
			ser, err := serial.New(serial.WithSerial(serNum))
			assert.NoError(t, err)
			w, err := New(WithSerial(ser))

			assert.NoError(t, err)
			assert.NoError(t, w.AddWire(tt.w1))
			assert.NoError(t, w.AddWire(tt.w2))
			assert.NoError(t, w.AddWire(tt.w3))
			assert.NoError(t, w.AddWire(tt.w4))
			assert.NoError(t, w.AddWire(tt.w5))
			s, err := w.Solve()
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, s)
		})
	}
}

func Test6(t *testing.T) {
	tests := []struct {
		name                   string
		w1, w2, w3, w4, w5, w6 string
		serial                 string
		expect                 string
	}{
		{
			name: "kein gelb + ungerade -> dritter draht",
			w1:   "B", w2: "B", w3: "W", w4: "W", w5: "S", w6: "W",
			serial: "gfsgfds1",
			expect: "Dritten Draht durchtrennen (W)",
		},
		{
			name: "ein gelb und mehr als ein weiß -> vierter draht",
			w1:   "W", w2: "G", w3: "W", w4: "B", w5: "S", w6: "S",
			serial: "gfsgfds2",
			expect: "Vierten Draht durchtrennen (B)",
		},
		{
			name: "kein rot -> letzter Draht",
			w1:   "B", w2: "B", w3: "W", w4: "W", w5: "S", w6: "W",
			serial: "gfsgfds2",
			expect: "Letzten Draht durchtrennen (W)",
		},
		{
			name: "ansonsten -> vierter draht",
			w1:   "B", w2: "B", w3: "W", w4: "W", w5: "S", w6: "R",
			serial: "gfsgfds2",
			expect: "Vierten Draht durchtrennen (W)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serNum := tt.serial
			if serNum == "" {
				serNum = "2"
			}
			ser, err := serial.New(serial.WithSerial(serNum))
			assert.NoError(t, err)
			w, err := New(WithSerial(ser))

			assert.NoError(t, err)
			assert.NoError(t, w.AddWire(tt.w1))
			assert.NoError(t, w.AddWire(tt.w2))
			assert.NoError(t, w.AddWire(tt.w3))
			assert.NoError(t, w.AddWire(tt.w4))
			assert.NoError(t, w.AddWire(tt.w5))
			assert.NoError(t, w.AddWire(tt.w6))
			s, err := w.Solve()
			assert.NoError(t, err)
			assert.Equal(t, tt.expect, s)
		})
	}
}
