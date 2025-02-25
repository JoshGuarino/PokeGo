package games

import (
	"github.com/JoshGuarino/PokeGo/internal/request"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

// Games group interface
type IGames interface {
}

// Games group struct
type Games struct{}

// Return an instance of Games resource group struct
func NewGamesGroup() Games {
	return Games{}
}

func GetGeneration() {

}

func GetGenerationList() {

}
