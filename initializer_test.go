package initializer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializer(t *testing.T) {
	c := 0
	init1 := func() error {
		c = c + 1
		return nil
	}
	init2 := func() error {
		c = c * 2
		return nil
	}

	Register("init1", init1)
	Register("init2", init2)

	assert.NotPanics(t, Init)
	assert.Equal(t, 2, c)
}

func TestInitializer_Error(t *testing.T) {
	c := 0
	init1 := func() error {
		c = c + 1
		return nil
	}
	init2 := func() error {
		c = c * 2
		return nil
	}
	init3 := func() error {
		return errors.New("testErr")
	}

	Register("init1", init1)
	Register("init2", init2)
	Register("init3", init3)

	assert.Panics(t, Init)
}
