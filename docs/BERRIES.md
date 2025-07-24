# [Berries Group](https://pokeapi.co/docs/v2#berries-section)

The following are the resources available in the Berries group of the [PokeAPI](https://pokeapi.co/).

## [Berry](https://pokeapi.co/docs/v2#berries)

##### Get single Berry resource by name or ID:

```go
// Main client example with ID
berry, err := client.Berries.GetBerry("1")

// Individual resource group example with Name
berry, err := berriesGroup.GetBerry("cheri")
```

##### Get list of Berry resource:

```go
// Main client example
berryList, err := client.Berries.GetBerryList(limit, offset)

// Individual resource group example 
berryList, err := berriesGroup.GetBerryList(limit, offset)
```

##### Get URL for Berry resource:
```go
// Main client example
berryURL, err := client.Berries.GetBerryURL()

// Individual resource group example
berryURL, err := berriesGroup.GetBerryURL()
```

## [BerryFirmness](https://pokeapi.co/docs/v2#berry-firmnesses)

##### Get single Berry Firmness resource by name or ID:

```go
// Main client example with Name
berryFirmness, err := client.Berries.GetBerryFirmness("soft")

// Individual resource group example with ID
berryFirmness, err := berriesGroup.GetBerryFirmness("1")
```

##### Get list of Berry Firmness resource:

```go
// Main client example returning 
berryFirmnessList, err := client.Berries.GetBerryFirmnessList(limit, offset)

// Individual resource group example 
berryFirmnessList, err := berriesGroup.GetBerryFirmnessList(limit, offset)
```

##### Get URL for Berry Firmness resource:
```go
// Main client example
berryFirmnessURL, err := client.Berries.GetBerryFirmnessURL()

// Individual resource group example
berryFirmnessURL, err := berriesGroup.GetBerryFirmnessURL()
```

## [BerryFlavor](https://pokeapi.co/docs/v2#berry-flavors)

##### Get single Berry Flavor resource by name or ID:

```go
// Main client example with ID
berryFlavor, err := client.Berries.GetBerryFlavor("spicy")

// Individual resource group example with Name
berryFlavor, err := berriesGroup.GetBerryFlavor("1")
```

##### Get list of Berry Flavor resource:

```go
// Main client example 
berryFlavorList, err := client.Berries.GetBerryFlavorList(limit, offset)

// Main client example 
berryFlavorList, err := berriesGroup.GetBerryFlavorList(limit, offset)
```

##### Get URL for Berry Flavor resource:
```go
// Main client example
berryFlavorURL, err := client.Berries.GetBerryFlavorURL()

// Individual resource group example
berryFlavorURL, err := berriesGroup.GetBerryFlavorURL()
```
