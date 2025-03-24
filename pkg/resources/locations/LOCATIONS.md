# [Locations Group](https://pokeapi.co/docs/v2#locations-section)

The following are the resources available in the Locations group of the [PokeAPI](https://pokeapi.co/).

## [Location](https://pokeapi.co/docs/v2#location)

##### Get single Location resource by Name or ID:

```go
// Main client example with ID
location, err := client.Locations.GetLocation("1")

// Individual resource group example with Name
location, err := locationsGroup.GetLocation("canalave-city")
```

##### Get list of Location resource:

```go
// Main client example returning first page of 20 results
locationList, err := client.Locations.GetLocationList(20, 0)

// Individual resource group example returning second page of 20 results
locationList, err := locationsGroup.GetLocationList(20, 20)
```

## [Location Area](https://pokeapi.co/docs/v2#location-areas)

##### Get single Location Area resource by Name or ID:

```go
// Main client example with ID
locationArea, err := client.Locations.GetLocationArea("1")

// Individual resource group example with Name
locationArea, err := locationsGroup.GetLocationArea("canalave-city-area")
```

##### Get list of Location Area resource:

```go
// Main client example returning first page of 20 results
locationAreaList, err := client.Locations.GetLocationAreaList(20, 0)

// Individual resource group example returning second page of 20 results
locationAreaList, err := locationsGroup.GetLocationAreaList(20, 20)
```

## [Pal Park Area](https://pokeapi.co/docs/v2#pal-park-areas)

##### Get single Pal Park Area resource by Name or ID:

```go
// Main client example with ID
palParkArea, err := client.Locations.GetPalParkArea("1")

// Individual resource group example with Name
palParkArea, err := locationsGroup.GetPalParkArea("forest")
```

##### Get list of Pal Park Area resource:

```go
// Main client example returning first page of 20 results
palParkAreaList, err := client.Locations.GetPalParkAreaList(20, 0)

// Individual resource group example returning second page of 20 results
palParkAreaList, err := locationsGroup.GetPalParkAreaList(20, 20)
```

## [Region](https://pokeapi.co/docs/v2#regions)

##### Get single Region resource by Name or ID:

```go
// Main client example with ID
region, err := client.Locations.GetRegion("1")

// Individual resource group example with Name
region, err := locationsGroup.GetRegion("kanto")
```

##### Get list of Region resource:

```go
// Main client example returning first page of 20 results
regionList, err := client.Locations.GetRegionList(20, 0)

// Individual resource group example returning second page of 20 results
regionList, err := locationsGroup.GetRegionList(20, 20)
```
