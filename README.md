# PokeGo <a href="https://pokeapi.co/api/v2/pokemon/charmander"><img src='https://veekun.com/dex/media/pokemon/global-link/4.png' height=100px/></a>

[![actions](https://github.com/JoshGuarino/PokeGo/actions/workflows/main.yml/badge.svg)](https://github.com/JoshGuarino/PokeGo/actions/workflows/main.yml)
![Release](https://img.shields.io/github/v/release/JoshGuarino/PokeGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoshGuarino/PokeGo?style=flat)](https://goreportcard.com/report/github.com/JoshGuarino/PokeGo)
[![license](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg "license")](https://github.com/JoshGuarino/PokeGo/blob/main/LICENSE)

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
limit, offset := 20, 0
pokemonList, err := client.Pokemon.GetPokemonList(limit, offset)

// Individual resource group example returning the second page of 20 results
limit, offset := 20, 20 
pokemonList, err := pokemonGroup.GetPokemonList(limit, offset)
```

## Caching

PokeGo uses a simple in-memory cache to store API responses. This is to reduce the number of requests made to the PokeAPI.
The cache is set to expire after 24 hours by default, as resources in the PokeAPI are mostly static.
You are able to disable the cache or change the expiration time of a cached resource.
I would reccommend against disabling it as it will result in a large number of requests to the PokeAPI and may result in rate limiting.
For more information on the cache, see the [CACHE.md](docs/CACHE.md) documentation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any suggestions or improvements.
See the [CONTRIBUTING.md](docs/CONTRIBUTING.md) file for more information.

## License

PokeGo is licensed under the [BSD 3-Clause License](LICENSE).
