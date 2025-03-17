# PokeGo <a href="https://pokeapi.co/api/v2/pokemon/charmander"><img src='https://veekun.com/dex/media/pokemon/global-link/4.png' height=50px/></a>

PokeGo is an Golang API wrapper for the [PokéAPI](https://pokeapi.co/) v2.

## Setup
There are 2 options for using PokeGo. You can either use the main client or create individual resource groups seperately. 
The main client will initialize all resource groups for you. If you choose to use individual resource groups, 
you will need to initialize each group separately.


##### Main client setup
```go 
import (
    pokego "github.com/JoshGuarino/PokeGo/pkg"
)

client := pokego.NewClient()
```
##### Individual resource group setup
```go
import (
    "github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

pokemonGroup := pokemon.NewPokemonGroup()
```

## Resource Groups
Below is a list of all the resource groups available in PokeGo. Each resource group has a set of methods that can be used to interact with the PokeAPI.
| Group | README | PokeAPI Docs |
| - | - | - |
| Berries | [README](pkg/resources/berries/README.md). | [Docs](https://pokeapi.co/docs/v2#berries-section) |
| Contests | [README](pkg/resources/contests/README.md). | [Docs](https://pokeapi.co/docs/v2#contests-section) |
| Encounters | [README](pkg/resources/encounters/README.md). | [Docs](https://pokeapi.co/docs/v2#encounters-section) |
| Evolution | [README](pkg/resources/evolution/README.md). | [Docs](https://pokeapi.co/docs/v2#evolution-section) |
| Games | [README](pkg/resources/games/README.md). | [Docs](https://pokeapi.co/docs/v2#games-section) |
| Items | [README](pkg/resources/items/README.md). | [Docs](https://pokeapi.co/docs/v2#items-section) |
| Locations | [README](pkg/resources/locations/README.md). | [Docs](https://pokeapi.co/docs/v2#locations-section) |
| Machines | [README](pkg/resources/machines/README.md). | [Docs](https://pokeapi.co/docs/v2#machines-section) |
| Moves | [README](pkg/resources/moves/README.md). | [Docs](https://pokeapi.co/docs/v2#moves-section) |
| Pokemon | [README](pkg/resources/pokemon/README.md). | [Docs](https://pokeapi.co/docs/v2#pokemon-section) |
| Utility | [README](pkg/resources/utility/README.md). | [Docs](https://pokeapi.co/docs/v2#utility-section) |

## Pagination
PokeGo uses the PokeAPI's pagination system to limit the number of results returned in a single request.
The default limit is 20 results per page and the default offset is 0. This will return the first 20 results.

## Caching
PokeGo uses a simple in-memory cache to store API responses. This is to reduce the number of requests made to the PokeAPI. 
The cache is set to expire after 24 hours as resources in the PokeAPI are mostly static. 
The cache is initialized when the client is created or when a resource group is created.
Only one instance of the cache is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new cache but reference the existing cache.
A pointer reference to the initialized cache is stored in the client and each resource group.

<details>
<summary>Clear Cache</summary>

##### The cache can be cleared by calling the `Clear()` method on the cache. 
```go
// Main client example
client.Cache.Clear()

// Individual resource group example
resourceGroup.Cache.Clear()
```
</details>

<details>
<summary>Disabling Cache</summary>

##### The active status of the cache can be set by calling the `setActive()` method on the cache. 
```go
// Main client example
client.Cache.SetActive(false)

// Individual resource group example
resourceGroup.Cache.SetActive(false)
```
</details>

<details>
<summary>Cache Status</summary>

##### The active status of the cache can be checked by calling the `GetActive()` method on the cache. 
```go
// Main client example
client.Cache.GetActive()

// Individual resource group example
resourceGroup.Cache.GetActive()
```
</details>

<details>
    <summary>Custom Expiration Time</summary>

##### The expiration time of the cache can be set by calling the `SetExpiration()` method on the cache. 
```go
// Main client example
client.Cache.SetExpiration(48 * time.Hour)

// Individual resource group example
resourceGroup.Cache.SetExpiration(48 * time.Hour)
```
</details>
