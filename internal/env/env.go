package env

import (
	"testing"

	"github.com/charmbracelet/log"
)

const (
	// Environment domains
	prodDomain  = "prod"
	stageDomain = "stage"

	// Base URLs
	prodURL  = "https://pokeapi.co/api/v2"
	stageURL = "https://staging.pokeapi.co/api/v2"
)

// Environment interface
type IEnv interface {
	GetDomain() string
	GetURL() string
	SetStage()
	SetProd()
}

// Environment struct
type Env struct {
	url    string
	domain string
}

// Environment global variable defaulting to production
var ENV *Env = prodEnv()

// Initialize function
func init() {
	// Set environment to stage if testing
	if testing.Testing() {
		ENV = stageEnv()
	}
	log.Info("Environment initialized", "domain", ENV.Domain(), "url", ENV.URL())
}

// Return an instance of Env struct
// pointing to the staging environment
func stageEnv() *Env {
	return &Env{
		url:    stageURL,
		domain: stageDomain,
	}
}

// Return an instance of Env struct
// pointing to the production environment
func prodEnv() *Env {
	return &Env{
		url:    prodURL,
		domain: prodDomain,
	}
}

// Return the environment domain
func (e *Env) Domain() string {
	return e.domain
}

// Return the url for the environment
func (e *Env) URL() string {
	return e.url
}

// Set the environment to stage
func (e *Env) SetStage() {
	e.domain = stageDomain
	e.url = stageURL
}

// Set the environment to production
func (e *Env) SetProd() {
	e.domain = prodDomain
	e.url = prodURL
}
