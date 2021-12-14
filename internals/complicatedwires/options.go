package complicatedwires

import "github.com/jcheyer/keep-talking-helper/internals/deps"

type Option func(c *cw)

func WithBomb(b deps.Bomb) Option {
	return func(c *cw) {
		c.bomb = b
	}
}
