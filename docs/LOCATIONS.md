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
// Main client example 
locationList, err := client.Locations.GetLocationList(limit, offset)

// Individual resource group example 
locationList, err := locationsGroup.GetLocationList(limit, offset)
```

##### Get the Location resource URL:

```go
// Main client example
locationURL := client.Locations.GetLocationURL()

// Individual resource group example
locationURL := locationsGroup.GetLocationURL()
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
// Main client example 
locationAreaList, err := client.Locations.GetLocationAreaList(limit, offset)

// Individual resource group example 
locationAreaList, err := locationsGroup.GetLocationAreaList(limit, offset)
```

##### Get the Location Area resource URL:

```go
// Main client example
locationAreaURL := client.Locations.GetLocationAreaURL()

// Individual resource group example
locationAreaURL := locationsGroup.GetLocationAreaURL()
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
// Main client example 
palParkAreaList, err := client.Locations.GetPalParkAreaList(limit, offset)

// Individual resource group example 
palParkAreaList, err := locationsGroup.GetPalParkAreaList(limit, offset)
```

##### Get the Pal Park Area resource URL:

```go
// Main client example
palParkAreaURL := client.Locations.GetPalParkAreaURL()

// Individual resource group example
palParkAreaURL := locationsGroup.GetPalParkAreaURL()
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
// Main client example 
regionList, err := client.Locations.GetRegionList(limit, offset)

// Individual resource group example 
regionList, err := locationsGroup.GetRegionList(limit, offset)
```

##### Get the Region resource URL:

```go
// Main client example
regionURL := client.Locations.GetRegionURL()

// Individual resource group example
regionURL := locationsGroup.GetRegionURL()
```
