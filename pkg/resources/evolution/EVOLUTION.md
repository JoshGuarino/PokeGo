# [Evolution Group](https://pokeapi.co/docs/v2#evolution-section)

The following are the resources available in the Evolution group of the [PokeAPI](https://pokeapi.co/).

## [EvolutionChain](https://pokeapi.co/docs/v2#evolution-chains)

##### Get single Evolution Chain resource by ID
```go
// Main client example
evolutionChain, err := client.Evolutions.GetEvolutionChain("1")

// Individual resource group example
evolutionChain, err := evolutionsGroup.GetEvolutionChain("1")
```

##### Get list of Evolution Chain resource
```go
// Main client example with no pagination options
evolutionChainList, err := client.Evolutions.GetEvolutionChainList(models.PaginationOptions{})

// Individual resource group example with pagination options
evolutionChainList, err := evolutionsGroup.GetEvolutionChainList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [EvolutionTrigger](https://pokeapi.co/docs/v2#evolution-triggers)

##### Get single Evolution Trigger resource by name or ID
```go
// Main client example with ID
evolutionTrigger, err := client.Evolutions.GetEvolutionTrigger("1")

// Individual resource group example with Name
evolutionTrigger, err := evolutionsGroup.GetEvolutionTrigger("level-up")
```

##### Get list of Evolution Trigger resource
```go
// Main client example with no pagination options
evolutionTriggerList, err := client.Evolutions.GetEvolutionTriggerList(models.PaginationOptions{})

// Individual resource group example with pagination options
evolutionTriggerList, err := evolutionsGroup.GetEvolutionTriggerList(models.PaginationOptions{Limit: 20, Offset: 20})
```
