package serial

import (
	"testing"

	"github.com/jcheyer/keep-talking-helper/internals/deps"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	ser, err := New()
	assert.NoError(t, err)
	assert.Implements(t, (*deps.Serial)(nil), ser)
}
