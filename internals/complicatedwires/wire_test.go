package complicatedwires

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWire(t *testing.T) {
	w := wire{
		color: "R",
	}
	assert.Equal(t, "0R0", w.String())
	w.led = true
	assert.Equal(t, "1R0", w.String())
	w.star = true
	assert.Equal(t, "1R1", w.String())

	w.color = "RW"
	assert.Equal(t, "1R1", w.SimplifiedString())
	w.color = "BW"
	assert.Equal(t, "1B1", w.SimplifiedString())
	w.color = "W"
	assert.Equal(t, "1W1", w.SimplifiedString())

}
