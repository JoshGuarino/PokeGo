package pokego

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/endpoints"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	pokego := NewClient()
	assert.IsType(t, PokeGo{}, pokego, "Expected PokeGo instance to be returned")
}

func TestRoot(t *testing.T) {
	root, err := NewClient().Root()
	assert.NoError(t, err, "Expected error to nil")
	assert.IsType(t, &models.Root{}, root, "Expected Root instance to be returned")
}

func TestGetBaseURL(t *testing.T) {
	baseURL := NewClient().GetBaseURL()
	assert.Equal(t, endpoints.StageBaseURL, baseURL, "Unexpected staging base URL returned")
	assert.IsType(t, "", baseURL, "Expected base URL to be a string")
}
