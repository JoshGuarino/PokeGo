# PokeGo

PokeGo is an Golang API wrapper for the [PokéAPI](https://pokeapi.co/) v2.

## Setup
There are 2 options for using PokeGo. You can either use the main client or create individual resource groups seperately. 
The main client will initialize all resource groups for you. If you choose to use individual resource groups, 
you will need to initialize each group separately.


Main client setup:
```go 
import (
    pokego "github.com/JoshGuarino/PokeGo/pkg"
)

client := pokego.NewClient()
```
Individual resource group setup:
```go
import (
    "github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

pokemonGroup := pokemon.NewPokemonGroup()
```

## Resources
Below is a list of all the resources available in PokeGo. Each resource group has a set of methods that can be used to interact with the PokeAPI.

### [Berries group](https://pokeapi.co/docs/v2#berries-section)
The berries resource group contains methods for interacting with the berries group resources. 
<details>
<summary>Berry</summary>
</details>

<details>
<summary>Berry Firmness</summary>
</details>

<details>
<summary>Berry Flavor</summary>
</details>

### [Contests group](https://pokeapi.co/docs/v2#contests-section)
The contests resource group contains methods for interacting with the contests group resources. 
<details>
<summary>Contest Types</summary>
</details>

<details>
<summary>Contest Effect</summary>
</details>

<details>
<summary>Super Contest Effect</summary>
</details>

### [Encounters group](https://pokeapi.co/docs/v2#encounters-section)
The encounters resource group contains methods for interacting with the Encounters group resources.
<details>
<summary>Encounter method</summary>
</details>

<details>
<summary>Encounter Condition</summary>
</details>

<details>
<summary>Encounter Condition Value</summary>
</details>

### [Evolution group](https://pokeapi.co/docs/v2#evolution-section)
The evolution resource group contains methods for interacting with the evolution group resources.
<details>
<summary>Evolution Chain</summary>
</details>

<details>
<summary>Evolution Trigger</summary>
</details>

### [Games group](https://pokeapi.co/docs/v2#games-section)
The games resource group contains methods for interacting with the  group resources.
<details>
<summary>Generation</summary>
</details>

<details>
<summary>Pokedex</summary>
</details>

<details>
<summary>Version</summary>
</details>

<details>
<summary>Version Group</summary>
</details>

### [Items group](https://pokeapi.co/docs/v2#items-section)
The items resource group contains methods for interacting with the items group resources.
<details>
<summary>Item</summary>
</details>

<details>
<summary>Item Attribute</summary>
</details>

<details>
<summary>Item Category</summary>
</details>

<details>
<summary>Item Fling Effect</summary>
</details>

<details>
<summary>Item Pocket</summary>
</details>

### [Locations group](https://pokeapi.co/docs/v2#locations-section)
The locations resource group contains methods for interacting with the Locations group resources.
<details>
<summary>Location</summary>
</details>

<details>
<summary>Location Area</summary>
</details>

<details>
<summary>Pal Park Area</summary>
</details>

<details>
<summary>Region</summary>
</details>

### [Machines group](https://pokeapi.co/docs/v2#machines-section)
The machines resource group contains methods for interacting with the machines group resources.
<details>
<summary>Machine</summary>
</details>

### [Moves group](https://pokeapi.co/docs/v2#moves-section)
The moves resource group contains methods for interacting with the moves group resources.
<details>
<summary>Move</summary>
</details>

<details>
<summary>Move Ailment</summary>
</details>

<details>
<summary>Move Battle Style</summary>
</details>

<details>
<summary>Move Category</summary>
</details>

<details>
<summary>Move Damage Class</summary>
</details>

<details>
<summary>Move Learn Method</summary>
</details>

<details>
<summary>Move Target</summary>
</details>

### [Pokemon group](https://pokeapi.co/docs/v2#pokemon-section)
The pokemon resource group contains methods for interacting with the pokemon group resources.
<details>
<summary>Ability</summary>
</details>

<details>
<summary>Characteristic</summary>
</details>

<details>
<summary>Egg Group</summary>
</details>

<details>
<summary>Gender</summary>
</details>

<details>
<summary>Growth Rate</summary>
</details>

<details>
<summary>Nature</summary>
</details>

<details>
<summary>Pokeathlon Stat</summary>
</details>

<details>
<summary>Pokemon</summary>
</details>

<details>
<summary>Pokemon Color</summary>
</details>

<details>
<summary>Pokemon Form</summary>
</details>

<details>
<summary>Pokemon Habitat</summary>
</details>

<details>
<summary>Pokemon Shape</summary>
</details>

<details>
<summary>Pokemon Species</summary>
</details>

<details>
<summary>Stat</summary>
</details>

<details>
<summary>Type</summary>
</details>

### [Utility group](https://pokeapi.co/docs/v2#utility-section)
The utility resource group contains methods for interacting with the utility group resources.
<details>
<summary>language</summary>
</details>

## Caching
PokeGo uses a simple in-memory cache to store API responses. This is to reduce the number of requests made to the PokeAPI. 
The cache is set to expire after 24 hours as resources in the PokeAPI are mostly static. 
The cache is initialized when the client is created or when a resource group is created.
Only one instance of the cache is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new cache but reference the existing cache.
A pointer reference to the initialized cache is stored in the client and each resource group.

<details>
<summary>Clear Cache</summary>

The cache can be cleared by calling the `Clear()` method on the cache. 
```go
// Main client example
client.Cache.Clear()

// Individual resource group example
resourceGroup.Cache.Clear()
```
</details>

<details>
<summary>Disabling Cache</summary>

The active status of the cache can be set by calling the `setActive()` method on the cache. 
```go
// Main client example
client.Cache.SetActive(false)

// Individual resource group example
resourceGroup.Cache.SetActive(false)
```
</details>

<details>
<summary>Cache Status</summary>

The active status of the cache can be checked by calling the `GetActive()` method on the cache. 
```go
// Main client example
client.Cache.GetActive()

// Individual resource group example
resourceGroup.Cache.GetActive()
```
</details>

<details>
    <summary>Custom Expiration Time</summary>

The expiration time of the cache can be set by calling the `SetExpiration()` method on the cache. 
```go
// Main client example
client.Cache.SetExpiration(48 * time.Hour)

// Individual resource group example
resourceGroup.Cache.SetExpiration(48 * time.Hour)
```
</details>
