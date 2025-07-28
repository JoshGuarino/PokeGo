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
// Main client example
generationList, err := client.Games.GetGenerationList(limit, offset)

// Individual resource group example 
generationList, err := gamesGroup.GetGenerationList(limit, offset)
```

##### Get Generation resource URL:

```go
// Main client example
generationURL := client.Games.GetGenerationURL()

// Individual resource group example
generationURL := gamesGroup.GetGenerationURL()
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
// Main client example 
pokedexList, err := client.Games.GetPokedexList(limit, offset)

// Individual resource group example
pokedexList, err := gamesGroup.GetPokedexList(limit, offset)
```

##### Get Pokedex resource URL:

```go
// Main client example
pokedexURL := client.Games.GetPokedexURL()

// Individual resource group example
pokedexURL := gamesGroup.GetPokedexURL()
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
// Main client example 
versionList, err := client.Games.GetVersionList(limit, offset)

// Individual resource group example 
versionList, err := gamesGroup.GetVersionList(limit, offset)
```

##### Get Version resource URL:

```go
// Main client example
versionURL := client.Games.GetVersionURL()

// Individual resource group example
versionURL := gamesGroup.GetVersionURL()
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
// Main client example 
versionGroupList, err := client.Games.GetVersionGroupList(limit, offset)

// Individual resource group example 
versionGroupList, err := gamesGroup.GetVersionGroupList(limit, offset)
```
