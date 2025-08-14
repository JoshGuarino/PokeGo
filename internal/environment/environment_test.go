package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var env *Environment = ENV

func TestStageEnv(t *testing.T) {
	stageEnv := stageEnv()
	assert.Equal(t, stageDomain, stageEnv.Domain(), "Unexpected environment domain")
	assert.Equal(t, stageURL, stageEnv.URL(), "Unexpected environment URL")
	assert.IsType(t, &Environment{}, stageEnv, "Expected environment struct instance to be returned")
}

func TestProdEnv(t *testing.T) {
	prodEnv := prodEnv()
	assert.Equal(t, prodDomain, prodEnv.Domain(), "Unexpected environment domain")
	assert.Equal(t, prodURL, prodEnv.URL(), "Unexpected environment URL")
	assert.IsType(t, &Environment{}, prodEnv, "Expected environment struct instance to be returned")
}

func TestGetDomain(t *testing.T) {
	assert.Equal(t, prodDomain, env.Domain(), "Unexpected environment domain")
	assert.IsType(t, "", env.Domain(), "Expected environment domain to be a string")
}

func TestGetURL(t *testing.T) {
	assert.Equal(t, prodURL, env.URL(), "Unexpected environment URL")
	assert.IsType(t, "", env.URL(), "Expected environment URL to be a string")
}

func TestSetStage(t *testing.T) {
	env.SetStage()
	assert.Equal(t, stageDomain, env.Domain(), "Unexpected environment domain")
	assert.Equal(t, stageURL, env.URL(), "Unexpected environment URL")
}

func TestSetProd(t *testing.T) {
	env.SetProd()
	assert.Equal(t, prodDomain, env.Domain(), "Unexpected environment domain")
	assert.Equal(t, prodURL, env.URL(), "Unexpected environment URL")
}
