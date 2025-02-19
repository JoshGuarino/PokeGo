package berries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var berries IBerries = Berries{}

func TestGetBerry(t *testing.T) {
	b, _ := berries.GetBerry("1")
	assert.Equal(t, "test", b.Name, "test")
}
