# [Evolution Group](https://pokeapi.co/docs/v2#evolution-section)

The following are the resources available in the Evolution group of the [PokeAPI](https://pokeapi.co/).

## [EvolutionChain](https://pokeapi.co/docs/v2#evolution-chains)

##### Get single Evolution Chain resource by ID:

```go
// Main client example
evolutionChain, err := client.Evolutions.GetEvolutionChain("1")

// Individual resource group example
evolutionChain, err := evolutionsGroup.GetEvolutionChain("1")
```

##### Get list of Evolution Chain resource:

```go
// Main client example returning first page of 20 results
evolutionChainList, err := client.Evolutions.GetEvolutionChainList(limit, offset)

// Individual resource group example returning second page of 20 results
evolutionChainList, err := evolutionsGroup.GetEvolutionChainList(limit, offset)
```

##### Get URL for Evolution Chain resource:
```go
// Main client example
evolutionChainURL := client.Evolutions.GetEvolutionChainURL()

// Individual resource group example
evolutionChainURL := evolutionsGroup.GetEvolutionChainURL()
```

## [EvolutionTrigger](https://pokeapi.co/docs/v2#evolution-triggers)

##### Get single Evolution Trigger resource by name or ID:

```go
// Main client example with ID
evolutionTrigger, err := client.Evolutions.GetEvolutionTrigger("1")

// Individual resource group example with Name
evolutionTrigger, err := evolutionsGroup.GetEvolutionTrigger("level-up")
```

##### Get list of Evolution Trigger resource:

```go
// Main client example returning first page of 20 results
evolutionTriggerList, err := client.Evolutions.GetEvolutionTriggerList(limit, offset)

// Individual resource group example returning second page of 20 results
evolutionTriggerList, err := evolutionsGroup.GetEvolutionTriggerList(limit, offset)
```

##### Get URL for Evolution Trigger resource:
```go
// Main client example
evolutionTriggerURL := client.Evolutions.GetEvolutionTriggerURL()

// Individual resource group example
evolutionTriggerURL := evolutionsGroup.GetEvolutionTriggerURL()
```
