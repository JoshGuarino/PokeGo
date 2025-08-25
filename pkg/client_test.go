package pokego

import (
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/environment"
	"github.com/JoshGuarino/PokeGo/pkg/models"
	"github.com/stretchr/testify/assert"
)

var pokego IPokeGo = NewClient()

func TestNewClient(t *testing.T) {
	assert.IsType(t, PokeGo{}, pokego, "Expected PokeGo instance to be returned")
}

func TestRoot(t *testing.T) {
	root, err := pokego.Root()
	assert.NoError(t, err, "Expected error to nil")
	assert.IsType(t, &models.Root{}, root, "Expected Root instance to be returned")
}

func TestGetBaseURL(t *testing.T) {
	baseURL := pokego.GetBaseURL()
	assert.Equal(t, environment.ENV.URL(), baseURL, "Unexpected staging base URL returned")
	assert.IsType(t, "", baseURL, "Expected base URL to be a string")
}
