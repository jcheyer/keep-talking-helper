package complicatedwires

import (
	"testing"

	"github.com/jcheyer/keep-talking-helper/internals/deps"
	"github.com/stretchr/testify/assert"
)

func TestAddWire(t *testing.T) {
	cw := New()
	err := cw.AddWire("gfdsgfds", false, false)
	assert.Error(t, err)
	err = cw.AddWire("b", false, false)
	assert.NoError(t, err)
	err = cw.AddWire("W", true, true)
	assert.NoError(t, err)
	assert.Len(t, cw.wires, 2)
	assert.Equal(t, "0B0", cw.wires[0].String())
	assert.Equal(t, "1W1", cw.wires[1].String())
}

func TestSolveWire(t *testing.T) {
	cw := New()
	w := wire{
		color: "BW",
	}
	cut, err := cw.SolveWire(w)
	assert.EqualError(t, err, deps.ErrSerNumUnknown.Error())
	assert.False(t, cut)

	err = cw.SetSerial("fdsa2")
	assert.NoError(t, err)

	cut, err = cw.SolveWire(w)
	assert.NoError(t, err)
	assert.True(t, cut)

	err = cw.SetSerial("fdsa1")
	assert.NoError(t, err)

	cut, err = cw.SolveWire(w)
	assert.NoError(t, err)
	assert.False(t, cut)

}

func TestSerSerial(t *testing.T) {
	cw := New()
	err := cw.SetSerial("")
	assert.Error(t, err)

	err = cw.SetSerial("hjk235")
	assert.NoError(t, err)
	even, _ := cw.serial.IsEven()
	assert.False(t, even)

	err = cw.SetSerial("hjk236")
	assert.NoError(t, err)
	even, _ = cw.serial.IsEven()
	assert.True(t, even)
}
