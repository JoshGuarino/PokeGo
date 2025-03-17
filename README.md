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
#### [Berries group](pkg/resources/berries/README.md)
#### [Contests Group](pkg/resources/contests/README.md)
#### [Encounters Group](pkg/resources/encounters/README.md)
#### [Evolution Group](pkg/resources/evolution/README.md)
#### [Games Group](pkg/resources/games/README.md)
#### [Items Group](pkg/resources/items/README.md)
#### [Locations Group](pkg/resources/locations/README.md)
#### [Machines Group](pkg/resources/machines/README.md)
#### [Moves Group](pkg/resources/moves/README.md)
#### [Pokemon Group](pkg/resources/pokemon/README.md)
#### [Utility Group](pkg/resources/utility/README.md)

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
