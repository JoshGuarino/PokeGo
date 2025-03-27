# [Games Group](https://pokeapi.co/docs/v2#games-section)

The following are the resources available in the Games group of the [PokeAPI](https://pokeapi.co/).

## [Generation](https://pokeapi.co/docs/v2#generations)

##### Get single Generation resource by Name or ID:

```go
// Main client example with ID
generation, err := client.Games.GetGeneration("1")

// Individual resource group example with Name
generation, err := gamesGroup.GetGeneration("generation-i")
```

##### Get list of Generation resource:

```go
// Main client example returning first page of 20 results
generationList, err := client.Games.GetGenerationList(20, 0)

// Individual resource group example returning second page of 20 results
generationList, err := gamesGroup.GetGenerationList(20, 20)
```

## [Pokedex](https://pokeapi.co/docs/v2#pokedexes)

##### Get single Pokedex resource by Name or ID:

```go
// Main client example with ID
pokedex, err := client.Games.GetPokedex("1")

// Individual resource group example with Name
pokedex, err := gamesGroup.GetPokedex("kanto")
```

##### Get list of Pokedex resource:

```go
// Main client example returning first page of 20 results
pokedexList, err := client.Games.GetPokedexList(20, 0)

// Individual resource group example returning second page of 20 results
pokedexList, err := gamesGroup.GetPokedexList(20, 20)
```

## [Version](https://pokeapi.co/docs/v2#versions)

##### Get single Version resource by Name or ID:

```go
// Main client example with ID
version, err := client.Games.GetVersion("1")

// Individual resource group example with Name
version, err := gamesGroup.GetVersion("red")
```

##### Get list of Version resource:

```go
// Main client example returning first page of 20 results
versionList, err := client.Games.GetVersionList(20, 0)

// Individual resource group example returning second page of 20 results
versionList, err := gamesGroup.GetVersionList(20, 20)
```

## [VersionGroup](https://pokeapi.co/docs/v2#version-groups)

##### Get single Version Group resource by Name or ID:

```go
// Main client example with ID
versionGroup, err := client.Games.GetVersionGroup("1")

// Individual resource group example with Name
versionGroup, err := gamesGroup.GetVersionGroup("red-blue")
```

##### Get list of Version Group resource:

```go
// Main client example returning first page of 20 results
versionGroupList, err := client.Games.GetVersionGroupList(20, 0)

// Individual resource group example returning second page of 20 results
versionGroupList, err := gamesGroup.GetVersionGroupList(20, 20)
```
