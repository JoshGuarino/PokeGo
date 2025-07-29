# [Pokemon Group](https://pokeapi.co/docs/v2#pokemon-section)

The following are the available resources in the Pokemon Group section of the [PokeAPI](https://pokeapi.co/).

## [Ability](https://pokeapi.co/docs/v2#abilities)

##### Get single Ability resource by name or ID:

```go
// Main client example with ID
ability, err := client.Pokemon.GetAbility("1")

// Individual resource group example with Name
ability, err := pokemonGroup.GetAbility("overgrow")
```

##### Get list of Ability resource:

```go
// Main client example returning first page of 20 results
abilityList, err := client.Pokemon.GetAbilityList(limit, offset)

// Individual resource group example returning second page of 20 results
abilityList, err := pokemonGroup.GetAbilityList(limit, offset)
```

##### Get Ability resource URL:

```go
// Main client example
abilityURL := client.Pokemon.GetAbilityURL()

// Individual resource group example
abilityURL := pokemonGroup.GetAbilityURL()
```

## [Characteristic](https://pokeapi.co/docs/v2#characteristics)

##### Get single Characteristic resource by ID:

```go
// Main client example with ID
characteristic, err := client.Pokemon.GetCharacteristic("1")

	models
// Individual resource group example with Name
characteristic, err := pokemonGroup.GetCharacteristic("1")
```

##### Get list of Characteristic resource:

```go
// Main client example returning first page of 20 results
characteristicList, err := client.Pokemon.GetCharacteristicList(limit, offset)

// Individual resource group example returning second page of 20 results
characteristicList, err := pokemonGroup.GetCharacteristicList(limit, offset)
```

##### Get Characteristic resource URL:

```go
// Main client example
characteristicURL := client.Pokemon.GetCharacteristicURL()

// Individual resource group example
characteristicURL := pokemonGroup.GetCharacteristicURL()
```

## [Egg Group](https://pokeapi.co/docs/v2#egg-groups)

##### Get single Egg Group resource by name or ID:

```go
// Main client example with ID
eggGroup, err := client.Pokemon.GetEggGroup("1")

// Individual resource group example with Name
eggGroup, err := pokemonGroup.GetEggGroup("monster")
```

##### Get list of Egg Group resource:

```go
// Main client example returning first page of 20 results
eggGroupList, err := client.Pokemon.GetEggGroupList(limit, offset)

// Individual resource group example returning second page of 20 results
eggGroupList, err := pokemonGroup.GetEggGroupList(limit, offset)
```

##### Get Egg Group resource URL:

```go
// Main client example
eggGroupURL := client.Pokemon.GetEggGroupURL()

// Individual resource group example
eggGroupURL := pokemonGroup.GetEggGroupURL()
```

## [Gender](https://pokeapi.co/docs/v2#gender)

##### Get single Gender resource by name or ID:

```go
// Main client example with ID
gender, err := client.Pokemon.GetGender("1")

// Individual resource group example with Name
gender, err := pokemonGroup.GetGender("male")
```

##### Get list of Gender resource:

```go
// Main client example returning first page of 20 results
genderList, err := client.Pokemon.GetGenderList(limit, offset)

// Individual resource group example returning second page of 20 results
genderList, err := pokemonGroup.GetGenderList(limit, offset)
```

##### Get Gender resource URL:

```go
// Main client example
genderURL := client.Pokemon.GetGenderURL()

// Individual resource group example
genderURL := pokemonGroup.GetGenderURL()
```

## [Growth Rate](https://pokeapi.co/docs/v2#growth-rates)

##### Get single Growth Rate resource by name or ID:

```go
// Main client example with ID
growthRate, err := client.Pokemon.GetGrowthRate("1")

// Individual resource group example with Name
growthRate, err := pokemonGroup.GetGrowthRate("slow")
```

##### Get list of Growth Rate resource:

```go
// Main client example returning first page of 20 results
growthRateList, err := client.Pokemon.GetGrowthRateList(limit, offset)

// Individual resource group example returning second page of 20 results
growthRateList, err := pokemonGroup.GetGrowthRateList(limit, offset)
```

##### Get Growth Rate resource URL:

```go
// Main client example
growthRateURL := client.Pokemon.GetGrowthRateURL()

// Individual resource group example
growthRateURL := pokemonGroup.GetGrowthRateURL()
```

## [Nature](https://pokeapi.co/docs/v2#natures)

##### Get single Nature resource by name or ID:

```go
// Main client example with ID
nature, err := client.Pokemon.GetNature("1")

// Individual resource group example with Name
nature, err := pokemonGroup.GetNature("hardy")
```

##### Get list of Nature resource:

```go
// Main client example returning first page of 20 results
natureList, err := client.Pokemon.GetNatureList(limit, offset)

// Individual resource group example returning second page of 20 results
natureList, err := pokemonGroup.GetNatureList(limit, offset)
```

##### Get Nature resource URL:

```go
// Main client example
natureURL := client.Pokemon.GetNatureURL()

// Individual resource group example
natureURL := pokemonGroup.GetNatureURL()
```

## [Pokeathlon Stat](https://pokeapi.co/docs/v2#pokeathlon-stats)

##### Get single Pokeathlon Stat resource by name or ID:

```go
// Main client example with ID
pokeathlonStat, err := client.Pokemon.GetPokeathlonStat("1")

// Individual resource group example with Name
pokeathlonStat, err := pokemonGroup.GetPokeathlonStat("speed")
```

##### Get list of Pokeathlon Stat resource:

```go
// Main client example returning first page of 20 results
pokeathlonStatList, err := client.Pokemon.GetPokeathlonStatList(limit, offset)

// Individual resource group example returning second page of 20 results
pokeathlonStatList, err := pokemonGroup.GetPokeathlonStatList(limit, offset)
```

##### Get Pokeathlon Stat resource URL:

```go
// Main client example
pokeathlonStatURL := client.Pokemon.GetPokeathlonStatURL()

// Individual resource group example
pokeathlonStatURL := pokemonGroup.GetPokeathlonStatURL()
```

## [Pokemon](https://pokeapi.co/docs/v2#pokemon)

##### Get single Pokemon resource by name or ID:

```go
// Main client example with ID
pokemon, err := client.Pokemon.GetPokemon("1")

// Individual resource group example with Name
pokemon, err := pokemonGroup.GetPokemon("bulbasaur")
```

##### Get list of Pokemon resource:

```go
// Main client example returning first page of 20 results
pokemonList, err := client.Pokemon.GetPokemonList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonList, err := pokemonGroup.GetPokemonList(limit, offset)
```

##### Get Pokemon resource URL:

```go
// Main client example
pokemonURL := client.Pokemon.GetPokemonURL()

// Individual resource group example
pokemonURL := pokemonGroup.GetPokemonURL()
```

## [Pokemon Location Areas](https://pokeapi.co/docs/v2#pokemon-location-areas)

##### Get list of Pokemon Location Areas:

```go
// Main client example with ID
pokemonLocationAreas, err := clinet.Pokemon.GetPokemonLocationAreas("1")

// Individual resource group example with Name
pokemonLocationAreas, err := pokemonGroup.GetPokemonLocationAreas("bulbasaur")
```

## [Pokemon Color](https://pokeapi.co/docs/v2#pokemon-colors)

##### Get single Pokemon Color resource by name or ID:

```go
// Main client example with ID
pokemonColor, err := client.Pokemon.GetPokemonColor("1")

// Individual resource group example with Name
pokemonColor, err := pokemonGroup.GetPokemonColor("black")
```

##### Get list of Pokemon Color resource:

```go
// Main client example returning first page of 20 results
pokemonColorList, err := client.Pokemon.GetPokemonColorList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonColorList, err := pokemonGroup.GetPokemonColorList(limit, offset)
```

##### Get Pokemon Color resource URL:

```go
// Main client example
pokemonColorURL := client.Pokemon.GetPokemonColorURL()

// Individual resource group example
pokemonColorURL := pokemonGroup.GetPokemonColorURL()
```

## [Pokemon Form](https://pokeapi.co/docs/v2#pokemon-forms)

##### Get single Pokemon Form resource by name or ID:

```go
// Main client example with ID
pokemonForm, err := client.Pokemon.GetPokemonForm("1")

// Individual resource group example with Name
pokemonForm, err := pokemonGroup.GetPokemonForm("bulbasaur")
```

##### Get list of Pokemon Form resource:

```go
// Main client example returning first page of 20 results
pokemonFormList, err := client.Pokemon.GetPokemonFormList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonFormList, err := pokemonGroup.GetPokemonFormList(limit, offset)
```

##### Get Pokemon Form resource URL:

```go
// Main client example
pokemonFormURL := client.Pokemon.GetPokemonFormURL()

// Individual resource group example
pokemonFormURL := pokemonGroup.GetPokemonFormURL()
```

## [Pokemon Habitat](https://pokeapi.co/docs/v2#pokemon-habitats)

##### Get single Pokemon Habitat resource by name or ID:

```go
// Main client example with ID
pokemonHabitat, err := client.Pokemon.GetPokemonHabitat("1")

// Individual resource group example with Name
pokemonHabitat, err := pokemonGroup.GetPokemonHabitat("cave")
```

##### Get list of Pokemon Habitat resource:

```go
// Main client example returning first page of 20 results
pokemonHabitatList, err := client.Pokemon.GetPokemonHabitatList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonHabitatList, err := pokemonGroup.GetPokemonHabitatList(limit, offset)
```

##### Get Pokemon Habitat resource URL:

```go
// Main client example
pokemonHabitatURL := client.Pokemon.GetPokemonHabitatURL()

// Individual resource group example
pokemonHabitatURL := pokemonGroup.GetPokemonHabitatURL()
```

## [Pokemon Shape](https://pokeapi.co/docs/v2#pokemon-shapes)

##### Get single Pokemon Shape resource by name or ID:

```go
// Main client example with ID
pokemonShape, err := client.Pokemon.GetPokemonShape("1")

// Individual resource group example with Name
pokemonShape, err := pokemonGroup.GetPokemonShape("ball")
```

##### Get list of Pokemon Shape resource:

```go
// Main client example returning first page of 20 results
pokemonShapeList, err := client.Pokemon.GetPokemonShapeList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonShapeList, err := pokemonGroup.GetPokemonShapeList(limit, offset)
```

##### Get Pokemon Shape resource URL:

```go
// Main client example
pokemonShapeURL := client.Pokemon.GetPokemonShapeURL()

// Individual resource group example
pokemonShapeURL := pokemonGroup.GetPokemonShapeURL()
```

## [Pokemon Species](https://pokeapi.co/docs/v2#pokemon-species)

##### Get single Pokemon Species resource by name or ID:

```go
// Main client example with ID
pokemonSpecies, err := client.Pokemon.GetPokemonSpecies("1")

// Individual resource group example with Name
pokemonSpecies, err := pokemonGroup.GetPokemonSpecies("bulbasaur")
```

##### Get list of Pokemon Species resource:

```go
// Main client example returning first page of 20 results
pokemonSpeciesList, err := client.Pokemon.GetPokemonSpeciesList(limit, offset)

// Individual resource group example returning second page of 20 results
pokemonSpeciesList, err := pokemonGroup.GetPokemonSpeciesList(limit, offset)
```

##### Get Pokemon Species resource URL:

```go
// Main client example
pokemonSpeciesURL := client.Pokemon.GetPokemonSpeciesURL()

// Individual resource group example
pokemonSpeciesURL := pokemonGroup.GetPokemonSpeciesURL()
```

## [Stat](https://pokeapi.co/docs/v2#stats)

##### Get single Stat resource by name or ID:

```go
// Main client example with ID
stat, err := client.Pokemon.GetStat("1")

// Individual resource group example with Name
stat, err := pokemonGroup.GetStat("hp")
```

##### Get list of Stat resource:

```go
// Main client example returning first page of 20 results
statList, err := client.Pokemon.GetStatList(limit, offset)

// Individual resource group example returning second page of 20 results
statList, err := pokemonGroup.GetStatList(limit, offset)
```

##### Get Stat resource URL:

```go
// Main client example
statURL := client.Pokemon.GetStatURL()

// Individual resource group example
statURL := pokemonGroup.GetStatURL()
```

## [Type](https://pokeapi.co/docs/v2#types)

##### Get single Type resource by name or ID:

```go
// Main client example with ID
typeResource, err := client.Pokemon.GetType("1")

// Individual resource group example with Name
typeResource, err := pokemonGroup.GetType("normal")
```

##### Get list of Type resource:

```go
// Main client example returning first page of 20 results
typeList, err := client.Pokemon.GetTypeList(limit, offset)

// Individual resource group example returning second page of 20 results
typeList, err := pokemonGroup.GetTypeList(limit, offset)
```

##### Get Type resource URL:

```go
// Main client example
typeURL := client.Pokemon.GetTypeURL()

// Individual resource group example
typeURL := pokemonGroup.GetTypeURL()
```
