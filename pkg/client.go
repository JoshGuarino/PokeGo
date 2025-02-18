package pokego

import (
	"github.com/JoshGuarino/PokeGo/internal/constants"
	"github.com/JoshGuarino/PokeGo/pkg/resources/berries"
)

// Return an instance of the PokeGo api wrapper client
func NewClient() PokeGo {
	return PokeGo{
		BaseUrl: constants.BaseUrl,
		Berries: berries.NewBerriesGroup(),
	}
}
