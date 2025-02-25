package pokego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	pokego := NewClient()
	assert.IsType(t, PokeGo{}, pokego, "Expected PokeGo instance to be returned")
}

func TestRoot(t *testing.T) {
	_, err := NewClient().Root()
	assert.NoError(t, err, "Expected error to nil")
}
