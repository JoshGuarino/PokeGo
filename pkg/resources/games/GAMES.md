# [Games Group](https://pokeapi.co/docs/v2#games-section)
The following are the resources available in the Games group of the [PokeAPI](https://pokeapi.co/).

## [Generation](https://pokeapi.co/docs/v2#generations)

##### Get single Generation resource by ID
```go
// Main client example with ID
generation, err := client.Games.GetGeneration("1")

// Individual resource group example with Name
generation, err := gamesGroup.GetGeneration("generation-i")
```

##### Get list of Generation resource
```go
// Main client example with no pagination options
generationList, err := client.Games.GetGenerationList(models.PaginationOptions{})

// Individual resource group example with pagination options
generationList, err := gamesGroup.GetGenerationList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokedex](https://pokeapi.co/docs/v2#pokedexes)

##### Get single Pokedex resource by ID
```go
// Main client example with ID
pokedex, err := client.Games.GetPokedex("1")

// Individual resource group example with Name
pokedex, err := gamesGroup.GetPokedex("kanto")
```

##### Get list of Pokedex resource
```go
// Main client example with no pagination options
pokedexList, err := client.Games.GetPokedexList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokedexList, err := gamesGroup.GetPokedexList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Version](https://pokeapi.co/docs/v2#versions)

##### Get single Version resource by ID
```go
// Main client example with ID
version, err := client.Games.GetVersion("1")

// Individual resource group example with Name
version, err := gamesGroup.GetVersion("red")
```

##### Get list of Version resource
```go
// Main client example with no pagination options
versionList, err := client.Games.GetVersionList(models.PaginationOptions{})

// Individual resource group example with pagination options
versionList, err := gamesGroup.GetVersionList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [VersionGroup](https://pokeapi.co/docs/v2#version-groups)

##### Get single Version Group resource by ID
```go
// Main client example with ID
versionGroup, err := client.Games.GetVersionGroup("1")

// Individual resource group example with Name
versionGroup, err := gamesGroup.GetVersionGroup("red-blue")
```

##### Get list of Version Group resource
```go
// Main client example with no pagination options
versionGroupList, err := client.Games.GetVersionGroupList(models.PaginationOptions{})

// Individual resource group example with pagination options
versionGroupList, err := gamesGroup.GetVersionGroupList(models.PaginationOptions{Limit: 20, Offset: 20})
```
