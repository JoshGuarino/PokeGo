# PokeGo <a href="https://pokeapi.co/api/v2/pokemon/charmander"><img src='https://veekun.com/dex/media/pokemon/global-link/4.png' height=100px/></a>

PokeGo is an Golang API wrapper for the [PokéAPI](https://pokeapi.co/) v2.

## Installation

PokeGo can be installed using the following command:

```bash
go get github.com/JoshGuarino/PokeGo
```

## Setup

There are two options for using PokeGo. You can either use the main client or create individual resource groups seperately.
The main client will initialize all resource groups for you. If you choose to use individual resource groups,
you will need to initialize each group separately.

##### Main client setup:

```go
import (
    pokego "github.com/JoshGuarino/PokeGo/pkg"
)

client := pokego.NewClient()
```

##### Individual resource group setup:

```go
import (
    "github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

pokemonGroup := pokemon.NewPokemonGroup()
```

## Resource Groups

Below is a list of all the resource groups available in PokeGo. Each resource group has a set of methods that can be used to interact with the PokeAPI.
| Group | Usage Docs | PokeAPI Docs |
| - | - | - |
| Berries | [BERRIES.md](docs/BERRIES.md) | https://pokeapi.co/docs/v2#berries-section |
| Contests | [CONTESTS.md](docs/CONTESTS.md) | https://pokeapi.co/docs/v2#contests-section |
| Encounters | [ENCOUNTERS.md](docs/ENCOUNTERS.md) | https://pokeapi.co/docs/v2#encounters-section |
| Evolution | [EVOLUTION.md](docs/EVOLUTION.md) | https://pokeapi.co/docs/v2#evolution-section |
| Games | [GAMES.md](docs/GAMES.md) | https://pokeapi.co/docs/v2#games-section |
| Items | [ITEMS.md](docs/ITEMS.md) | https://pokeapi.co/docs/v2#items-section |
| Locations | [LOCATIONS.md](docs/LOCATIONS.md) | https://pokeapi.co/docs/v2#locations-section |
| Machines | [MACHINES.md](docs/MACHINES.md) | https://pokeapi.co/docs/v2#machines-section |
| Moves | [MOVES.md](docs/MOVES.md) | https://pokeapi.co/docs/v2#moves-section |
| Pokemon | [POKEMON.md](docs/POKEMON.md) | https://pokeapi.co/docs/v2#pokemon-section |
| Utility | [UTILITY.md](docs/UTILITY.md) | https://pokeapi.co/docs/v2#utility-section |

## Pagination

PokeGo supports pagination for list endpoints. The `GetPokemonList()` method is an example of a list endpoint that supports pagination.
The method takes two arguments, `limit` and `offset`. The `limit` argument is the number of results to return and the `offset` argument
is the number of results to skip. Both arugments are required as Golang does not support default arguments.

##### List endpoint pagination examples:

```go
// Main client example returning the first page of 20 results
pokemonList, err := client.Pokemon.GetPokemonList(20, 0)

// Individual resource group example returning the second page of 20 results
pokemonList, err := pokemonGroup.GetPokemonList(20, 20)
```

## Caching

PokeGo uses a simple in-memory cache to store API responses. This is to reduce the number of requests made to the PokeAPI.
The cache is set to expire after 24 hours as resources in the PokeAPI are mostly static.
The cache is initialized when the client is created or when a resource group is created.
Only one instance of the cache is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new cache but reference the existing cache.
A pointer reference to the initialized cache is stored in the client and each resource group.
For more information on the cache, see the [cache](docs/CACHE.md) documentation.
