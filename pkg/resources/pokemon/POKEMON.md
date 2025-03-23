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
// Main client example returning first page of 20 results
abilityList, err := client.Pokemon.GetAbilityList(20, 0)

// Individual resource group example returning second page of 20 results
abilityList, err := pokemonGroup.GetAbilityList(20, 20)
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
// Main client example returning first page of 20 results
characteristicList, err := client.Pokemon.GetCharacteristicList(20, 0)

// Individual resource group example returning second page of 20 results
characteristicList, err := pokemonGroup.GetCharacteristicList(20, 20)
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
// Main client example returning first page of 20 results
eggGroupList, err := client.Pokemon.GetEggGroupList(20, 0)

// Individual resource group example returning second page of 20 results
eggGroupList, err := pokemonGroup.GetEggGroupList(20, 20)
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
// Main client example returning first page of 20 results
genderList, err := client.Pokemon.GetGenderList(20, 0)

// Individual resource group example returning second page of 20 results
genderList, err := pokemonGroup.GetGenderList(20, 20)
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
// Main client example returning first page of 20 results
growthRateList, err := client.Pokemon.GetGrowthRateList(20, 0)

// Individual resource group example returning second page of 20 results
growthRateList, err := pokemonGroup.GetGrowthRateList(20, 20)
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
// Main client example returning first page of 20 results
natureList, err := client.Pokemon.GetNatureList(20, 0)

// Individual resource group example returning second page of 20 results
natureList, err := pokemonGroup.GetNatureList(20, 20)
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
// Main client example returning first page of 20 results
pokeathlonStatList, err := client.Pokemon.GetPokeathlonStatList(20, 0)

// Individual resource group example returning second page of 20 results
pokeathlonStatList, err := pokemonGroup.GetPokeathlonStatList(20, 20)
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
// Main client example returning first page of 20 results
pokemonList, err := client.Pokemon.GetPokemonList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonList, err := pokemonGroup.GetPokemonList(20, 20)
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
// Main client example returning first page of 20 results
pokemonColorList, err := client.Pokemon.GetPokemonColorList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonColorList, err := pokemonGroup.GetPokemonColorList(20, 20)
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
// Main client example returning first page of 20 results
pokemonFormList, err := client.Pokemon.GetPokemonFormList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonFormList, err := pokemonGroup.GetPokemonFormList(20, 20)
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
// Main client example returning first page of 20 results
pokemonHabitatList, err := client.Pokemon.GetPokemonHabitatList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonHabitatList, err := pokemonGroup.GetPokemonHabitatList(20, 20)
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
// Main client example returning first page of 20 results
pokemonShapeList, err := client.Pokemon.GetPokemonShapeList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonShapeList, err := pokemonGroup.GetPokemonShapeList(20, 20)
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
// Main client example returning first page of 20 results
pokemonSpeciesList, err := client.Pokemon.GetPokemonSpeciesList(20, 0)

// Individual resource group example returning second page of 20 results
pokemonSpeciesList, err := pokemonGroup.GetPokemonSpeciesList(20, 20)
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
// Main client example returning first page of 20 results
statList, err := client.Pokemon.GetStatList(20, 0)

// Individual resource group example returning second page of 20 results
statList, err := pokemonGroup.GetStatList(20, 20)
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
// Main client example returning first page of 20 results
typeList, err := client.Pokemon.GetTypeList(20, 0)

// Individual resource group example returning second page of 20 results
typeList, err := pokemonGroup.GetTypeList(20, 20)
```
