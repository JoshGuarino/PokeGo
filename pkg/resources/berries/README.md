# Berries Group
The following are the resources available in the Berries group of the [PokeAPI](https://pokeapi.co/docs/v2#berries-section).

## Berry

##### Get single Berry resource by name or ID
```go
// Main client example with ID
berry, err := client.Berries.GetBerry("1")

// Individual resource group example with Name
berry, err := berriesGroup.GetBerry("cheri")
```

##### Get list of Berry resource
```go
// Main client example with no pagination options
berryList, err := client.Berries.GetBerryList(models.PaginationOptions{})

// Individual resource group example with pagination options
berryList, err := berriesGroup.GetBerryList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## BerryFirmness

##### Get single Berry Firmness resource by name or ID
```go
// Main client example with Name
berryFirmness, err := client.Berries.GetBerryFirmness("soft")

// Individual resource group example with ID
berryFirmness, err := berriesGroup.GetBerryFirmness("1")
```
##### Get list of Berry Firmness resource
```go
// Main client example with pagination options
berryFirmnessList, err := client.Berries.GetBerryFirmnessList(models.PaginationOptions{Limit: 20, Offset: 20})

// Individual resource group example without pagination options
berryFirmnessList, err := berriesGroup.GetBerryFirmnessList(models.PaginationOptions{})
```
## BerryFlavor

##### Get single Berry Flavor resource by name or ID
```go
// Main client example with ID
berryFlavor, err := client.Berries.GetBerryFlavor("spicy")

// Individual resource group example with Name
berryFlavor, err := berriesGroup.GetBerryFlavor("1")
```

##### Get list of Berry Flavor resource
```go
// Main client example with pagination options
berryFlavorList, err := client.Berries.GetBerryFlavorList(models.PaginationOptions{Limit: 20, Offset: 20})

// Individual resource group example without pagination options
berryFlavorList, err := berriesGroup.GetBerryFlavorList(models.PaginationOptions{})
```
