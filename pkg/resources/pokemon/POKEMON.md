# [Pokemon Group](https://pokeapi.co/docs/v2#pokemon-section)
The following are the available resources in the Pokemon Group section of the [PokeAPI](https://pokeapi.co/).

 ## [Ability](https://pokeapi.co/docs/v2#abilities)

##### Get single Ability resource by name or ID
```go
// Main client example with ID
ability, err := client.Pokemon.GetAbility("1")

// Individual resource group example with Name
ability, err := pokemonGroup.GetAbility("overgrow")
```

##### Get list of Ability resource
```go
// Main client example with no pagination options
abilityList, err := client.Pokemon.GetAbilityList(models.PaginationOptions{})

// Individual resource group example with pagination options
abilityList, err := pokemonGroup.GetAbilityList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Characteristic](https://pokeapi.co/docs/v2#characteristics)

##### Get single Characteristic resource by ID
```go
// Main client example with ID
characteristic, err := client.Pokemon.GetCharacteristic("1")

// Individual resource group example with Name
characteristic, err := pokemonGroup.GetCharacteristic("1")
```

##### Get list of Characteristic resource
```go
// Main client example with no pagination options
characteristicList, err := client.Pokemon.GetCharacteristicList(models.PaginationOptions{})

// Individual resource group example with pagination options
characteristicList, err := pokemonGroup.GetCharacteristicList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Egg Group](https://pokeapi.co/docs/v2#egg-groups)

##### Get single Egg Group resource by name or ID
```go
// Main client example with ID
eggGroup, err := client.Pokemon.GetEggGroup("1")

// Individual resource group example with Name
eggGroup, err := pokemonGroup.GetEggGroup("monster")
```

##### Get list of Egg Group resource
```go
// Main client example with no pagination options
eggGroupList, err := client.Pokemon.GetEggGroupList(models.PaginationOptions{})

// Individual resource group example with pagination options
eggGroupList, err := pokemonGroup.GetEggGroupList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Gender](https://pokeapi.co/docs/v2#gender)

##### Get single Gender resource by name or ID
```go
// Main client example with ID
gender, err := client.Pokemon.GetGender("1")

// Individual resource group example with Name
gender, err := pokemonGroup.GetGender("male")
```

##### Get list of Gender resource
```go
// Main client example with no pagination options
genderList, err := client.Pokemon.GetGenderList(models.PaginationOptions{})

// Individual resource group example with pagination options
genderList, err := pokemonGroup.GetGenderList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Growth Rate](https://pokeapi.co/docs/v2#growth-rates)

##### Get single Growth Rate resource by name or ID
```go
// Main client example with ID
growthRate, err := client.Pokemon.GetGrowthRate("1")

// Individual resource group example with Name
growthRate, err := pokemonGroup.GetGrowthRate("slow")
```

##### Get list of Growth Rate resource
```go
// Main client example with no pagination options
growthRateList, err := client.Pokemon.GetGrowthRateList(models.PaginationOptions{})

// Individual resource group example with pagination options
growthRateList, err := pokemonGroup.GetGrowthRateList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Nature](https://pokeapi.co/docs/v2#natures)

##### Get single Nature resource by name or ID
```go
// Main client example with ID
nature, err := client.Pokemon.GetNature("1")

// Individual resource group example with Name
nature, err := pokemonGroup.GetNature("hardy")
```

##### Get list of Nature resource
```go
// Main client example with no pagination options
natureList, err := client.Pokemon.GetNatureList(models.PaginationOptions{})

// Individual resource group example with pagination options
natureList, err := pokemonGroup.GetNatureList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokeathlon Stat](https://pokeapi.co/docs/v2#pokeathlon-stats)

##### Get single Pokeathlon Stat resource by name or ID
```go
// Main client example with ID
pokeathlonStat, err := client.Pokemon.GetPokeathlonStat("1")

// Individual resource group example with Name
pokeathlonStat, err := pokemonGroup.GetPokeathlonStat("speed")
```

##### Get list of Pokeathlon Stat resource
```go
// Main client example with no pagination options
pokeathlonStatList, err := client.Pokemon.GetPokeathlonStatList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokeathlonStatList, err := pokemonGroup.GetPokeathlonStatList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon](https://pokeapi.co/docs/v2#pokemon)

##### Get single Pokemon resource by name or ID
```go
// Main client example with ID
pokemon, err := client.Pokemon.GetPokemon("1")

// Individual resource group example with Name
pokemon, err := pokemonGroup.GetPokemon("bulbasaur")
```

##### Get list of Pokemon resource
```go
// Main client example with no pagination options
pokemonList, err := client.Pokemon.GetPokemonList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonList, err := pokemonGroup.GetPokemonList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon Color](https://pokeapi.co/docs/v2#pokemon-colors)

##### Get single Pokemon Color resource by name or ID
```go
// Main client example with ID
pokemonColor, err := client.Pokemon.GetPokemonColor("1")

// Individual resource group example with Name
pokemonColor, err := pokemonGroup.GetPokemonColor("black")
```

##### Get list of Pokemon Color resource
```go
// Main client example with no pagination options
pokemonColorList, err := client.Pokemon.GetPokemonColorList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonColorList, err := pokemonGroup.GetPokemonColorList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon Form](https://pokeapi.co/docs/v2#pokemon-forms)

##### Get single Pokemon Form resource by name or ID
```go
// Main client example with ID
pokemonForm, err := client.Pokemon.GetPokemonForm("1")

// Individual resource group example with Name
pokemonForm, err := pokemonGroup.GetPokemonForm("bulbasaur")
```

##### Get list of Pokemon Form resource
```go
// Main client example with no pagination options
pokemonFormList, err := client.Pokemon.GetPokemonFormList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonFormList, err := pokemonGroup.GetPokemonFormList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon Habitat](https://pokeapi.co/docs/v2#pokemon-habitats)

##### Get single Pokemon Habitat resource by name or ID
```go
// Main client example with ID
pokemonHabitat, err := client.Pokemon.GetPokemonHabitat("1")

// Individual resource group example with Name
pokemonHabitat, err := pokemonGroup.GetPokemonHabitat("cave")
```

##### Get list of Pokemon Habitat resource
```go
// Main client example with no pagination options
pokemonHabitatList, err := client.Pokemon.GetPokemonHabitatList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonHabitatList, err := pokemonGroup.GetPokemonHabitatList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon Shape](https://pokeapi.co/docs/v2#pokemon-shapes)

##### Get single Pokemon Shape resource by name or ID
```go
// Main client example with ID
pokemonShape, err := client.Pokemon.GetPokemonShape("1")

// Individual resource group example with Name
pokemonShape, err := pokemonGroup.GetPokemonShape("ball")
```

##### Get list of Pokemon Shape resource
```go
// Main client example with no pagination options
pokemonShapeList, err := client.Pokemon.GetPokemonShapeList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonShapeList, err := pokemonGroup.GetPokemonShapeList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Pokemon Species](https://pokeapi.co/docs/v2#pokemon-species)

##### Get single Pokemon Species resource by name or ID
```go
// Main client example with ID
pokemonSpecies, err := client.Pokemon.GetPokemonSpecies("1")

// Individual resource group example with Name
pokemonSpecies, err := pokemonGroup.GetPokemonSpecies("bulbasaur")
```

##### Get list of Pokemon Species resource
```go
// Main client example with no pagination options
pokemonSpeciesList, err := client.Pokemon.GetPokemonSpeciesList(models.PaginationOptions{})

// Individual resource group example with pagination options
pokemonSpeciesList, err := pokemonGroup.GetPokemonSpeciesList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Stat](https://pokeapi.co/docs/v2#stats)

##### Get single Stat resource by name or ID
```go
// Main client example with ID
stat, err := client.Pokemon.GetStat("1")

// Individual resource group example with Name
stat, err := pokemonGroup.GetStat("hp")
```

##### Get list of Stat resource
```go
// Main client example with no pagination options
statList, err := client.Pokemon.GetStatList(models.PaginationOptions{})

// Individual resource group example with pagination options
statList, err := pokemonGroup.GetStatList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Type](https://pokeapi.co/docs/v2#types)

##### Get single Type resource by name or ID
```go
// Main client example with ID
typeResource, err := client.Pokemon.GetType("1")

// Individual resource group example with Name
typeResource, err := pokemonGroup.GetType("normal")
```

##### Get list of Type resource
```go
// Main client example with no pagination options
typeList, err := client.Pokemon.GetTypeList(models.PaginationOptions{})

// Individual resource group example with pagination options
typeList, err := pokemonGroup.GetTypeList(models.PaginationOptions{Limit: 20, Offset: 20})
```
