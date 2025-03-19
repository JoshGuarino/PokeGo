# [Locations Group](https://pokeapi.co/docs/v2#locations-section)
The following are the resources available in the Locations group of the [PokeAPI](https://pokeapi.co/).

## [Location](https://pokeapi.co/docs/v2#location)

##### Get single Location resource by Name or ID
```go
// Main client example with ID
location, err := client.Locations.GetLocation("1")

// Individual resource group example with Name
location, err := locationsGroup.GetLocation("kanto")
```

##### Get list of Location resource 
```go
// Main client example with no pagination options
locationList, err := client.Locations.GetLocationList(models.PaginationOptions{})

// Individual resource group example with pagination options
locationList, err := locationsGroup.GetLocationList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Location Area](https://pokeapi.co/docs/v2#location-areas)

##### Get single Location Area resource by Name or ID
```go
// Main client example with ID
locationArea, err := client.Locations.GetLocationArea("1")

// Individual resource group example with Name
locationArea, err := locationsGroup.GetLocationArea("canalave-city-area")
```

##### Get list of Location Area resource 
```go
// Main client example with no pagination options
locationAreaList, err := client.Locations.GetLocationAreaList(models.PaginationOptions{})

// Individual resource group example with pagination options
locationAreaList, err := locationsGroup.GetLocationAreaList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pal Park Area](https://pokeapi.co/docs/v2#pal-park-areas)

##### Get single Pal Park Area resource by Name or ID
```go
// Main client example with ID
palParkArea, err := client.Locations.GetPalParkArea("1")

// Individual resource group example with Name
palParkArea, err := locationsGroup.GetPalParkArea("forest")
```

##### Get list of Pal Park Area resource 
```go
// Main client example with no pagination options
palParkAreaList, err := client.Locations.GetPalParkAreaList(models.PaginationOptions{})

// Individual resource group example with pagination options
palParkAreaList, err := locationsGroup.GetPalParkAreaList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Region](https://pokeapi.co/docs/v2#regions)

##### Get single Region resource by Name or ID
```go
// Main client example with ID
region, err := client.Locations.GetRegion("1")

// Individual resource group example with Name
region, err := locationsGroup.GetRegion("kanto")
```

##### Get list of Region resource 
```go
// Main client example with no pagination options
regionList, err := client.Locations.GetRegionList(models.PaginationOptions{})

// Individual resource group example with pagination options
regionList, err := locationsGroup.GetRegionList(models.PaginationOptions{Limit: 20, Offset: 20})
```
