package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const secret = "secret"

func TestState(t *testing.T) {
	state := Generate(secret)
	assert.NotNil(t, state)
	valid := Valid(state, secret)
	assert.True(t, valid)
	valid = Valid("", secret)
	assert.False(t, valid)
	valid = Valid(secret, secret)
	assert.False(t, valid)
	valid = Valid("a"+state, secret)
	assert.False(t, valid)
}
