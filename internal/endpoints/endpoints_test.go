package endpoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseURL(t *testing.T) {
	assert.Equal(t, StageBaseURL, BaseURL, "Expected BaseURL to be set to StageBaseURL")
}
