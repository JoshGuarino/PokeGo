# [Berries Group](https://pokeapi.co/docs/v2#berries-section)

The following are the resources available in the Berries group of the [PokeAPI](https://pokeapi.co/).

## [Berry](https://pokeapi.co/docs/v2#berries)

##### Get single Berry resource by name or ID

```go
// Main client example with ID
berry, err := client.Berries.GetBerry("1")

// Individual resource group example with Name
berry, err := berriesGroup.GetBerry("cheri")
```

##### Get list of Berry resource

```go
// Main client example returning first page of 20 results
berryList, err := client.Berries.GetBerryList(20, 0)

// Individual resource group example returning second page of 20 results
berryList, err := berriesGroup.GetBerryList(20, 20)
```

## [BerryFirmness](https://pokeapi.co/docs/v2#berry-firmnesses)

##### Get single Berry Firmness resource by name or ID

```go
// Main client example with Name
berryFirmness, err := client.Berries.GetBerryFirmness("soft")

// Individual resource group example with ID
berryFirmness, err := berriesGroup.GetBerryFirmness("1")
```

##### Get list of Berry Firmness resource

```go
// Main client example returning first page of 20 results
berryFirmnessList, err := client.Berries.GetBerryFirmnessList(20, 0)

// Individual resource group example returning second page of 20 results
berryFirmnessList, err := berriesGroup.GetBerryFirmnessList(20, 20)
```

## [BerryFlavor](https://pokeapi.co/docs/v2#berry-flavors)

##### Get single Berry Flavor resource by name or ID

```go
// Main client example with ID
berryFlavor, err := client.Berries.GetBerryFlavor("spicy")

// Individual resource group example with Name
berryFlavor, err := berriesGroup.GetBerryFlavor("1")
```

##### Get list of Berry Flavor resource

```go
// Main client example returning first page of 20 results
berryFlavorList, err := client.Berries.GetBerryFlavorList(20, 0)

// Main client example returning first page of 20 results
berryFlavorList, err := berriesGroup.GetBerryFlavorList(20, 20)
```
