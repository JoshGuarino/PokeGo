# [Contests Group](https://pokeapi.co/docs/v2#contests-section)

The following are the resources available in the Contests group of the [PokeAPI](https://pokeapi.co/).

## [ContestType](https://pokeapi.co/docs/v2#contest-types)

##### Get single Contest Type resource by name or ID

```go
// Main client example with ID
contestType, err := client.Contests.GetContestType("1")

// Individual resource group example with Name
contestType, err := contestsGroup.GetContestType("cool")
```

##### Get list of Contest Type resource

```go
// Main client example returning first page of 20 results
contestTypeList, err := client.Contests.GetContestTypeList(20, 0)

// Individual resource group example returning second page of 20 results
contestTypeList, err := contestsGroup.GetContestTypeList(20, 20)
```

## [ContestEffect](https://pokeapi.co/docs/v2#contest-effects)

##### Get single Contest Effect resource by ID

```go
// Main client example
contestEffect, err := client.Contests.GetContestEffect("1")

// Individual resource group example
contestEffect, err := contestsGroup.GetContestEffect("1")
```

##### Get list of Contest Effect resource

```go
// Main client example returning first page of 20 results
contestEffectList, err := client.Contests.GetContestEffectList(20, 0)

// Individual resource group example returning second page of 20 results
contestEffectList, err := contestsGroup.GetContestEffectList(20, 20)
```

## [SuperContestEffect](https://pokeapi.co/docs/v2#super-contest-effects)

##### Get single Super Contest Effect resource by ID

```go
// Main client example
superContestEffect, err := client.Contests.GetSuperContestEffect("1")

// Individual resource group example
superContestEffect, err := contestsGroup.GetSuperContestEffect("1")
```

##### Get list of Super Contest Effect resource

```go
// Main client example returning first page of 20 results
superContestEffectList, err := client.Contests.GetSuperContestEffectList(20, 0)

// Individual resource group example returning second page of 20 results
superContestEffectList, err := contestsGroup.GetSuperContestEffectList(20, 20)
```
