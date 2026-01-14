package environment

import (
	"sync"
	"testing"

	"github.com/JoshGuarino/PokeGo/internal/logger"
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
type IEnvironment interface {
	Domain() string
	URL() string
	SetStage()
	SetProd()
}

// Environment struct
type Environment struct {
	url    string
	domain string
	lock   sync.Mutex
}

// Environment global variable defaulting to production
var ENV *Environment = prodEnv()

// Logger reference variable
var log logger.ILogger = logger.LOG

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
func stageEnv() *Environment {
	return &Environment{
		url:    stageURL,
		domain: stageDomain,
	}
}

// Return an instance of Env struct
// pointing to the production environment
func prodEnv() *Environment {
	return &Environment{
		url:    prodURL,
		domain: prodDomain,
	}
}

// Return the environment domain
func (e *Environment) Domain() string {
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.domain
}

// Return the url for the environment
func (e *Environment) URL() string {
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.url
}

// Set the environment to stage
func (e *Environment) SetStage() {
	log.Warn("Environment changed to stage", "domain", stageDomain, "url", stageURL)
	e.lock.Lock()
	defer e.lock.Unlock()
	e.domain = stageDomain
	e.url = stageURL
}

// Set the environment to production
func (e *Environment) SetProd() {
	log.Warn("Environment changed to production", "domain", prodDomain, "url", prodURL)
	e.lock.Lock()
	defer e.lock.Unlock()
	e.domain = prodDomain
	e.url = prodURL
}
