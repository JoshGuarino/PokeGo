# [Contests Group](https://pokeapi.co/docs/v2#contests-section)

The following are the resources available in the Contests group of the [PokeAPI](https://pokeapi.co/).

## [ContestType](https://pokeapi.co/docs/v2#contest-types)

##### Get single Contest Type resource by name or ID:

```go
// Main client example with ID
contestType, err := client.Contests.GetContestType("1")

// Individual resource group example with Name
contestType, err := contestsGroup.GetContestType("cool")
```

##### Get list of Contest Type resource:

```go
// Main client example 
contestTypeList, err := client.Contests.GetContestTypeList(limit, offset)

// Individual resource group example 
contestTypeList, err := contestsGroup.GetContestTypeList(limit, offset)
```

##### Get URL for Contest Type resource:
```go
// Main client example
contestTypeURL := client.Contests.GetContestTypeURL()

// Individual resource group example
contestTypeURL := contestsGroup.GetContestTypeURL()
```

## [ContestEffect](https://pokeapi.co/docs/v2#contest-effects)

##### Get single Contest Effect resource by ID:

```go
// Main client example
contestEffect, err := client.Contests.GetContestEffect("1")

// Individual resource group example
contestEffect, err := contestsGroup.GetContestEffect("1")
```

##### Get list of Contest Effect resource:

```go
// Main client example 
contestEffectList, err := client.Contests.GetContestEffectList(limit, offset)

// Individual resource group example 
contestEffectList, err := contestsGroup.GetContestEffectList(limit, offset)
```

##### Get URL for Contest Effect resource:
```go
// Main client example
contestEffectURL := client.Contests.GetContestEffectURL()

// Individual resource group example
contestEffectURL := contestsGroup.GetContestEffectURL()
```

## [SuperContestEffect](https://pokeapi.co/docs/v2#super-contest-effects)

##### Get single Super Contest Effect resource by ID:

```go
// Main client example
superContestEffect, err := client.Contests.GetSuperContestEffect("1")

// Individual resource group example
superContestEffect, err := contestsGroup.GetSuperContestEffect("1")
```

##### Get list of Super Contest Effect resource:

```go
// Main client example 
superContestEffectList, err := client.Contests.GetSuperContestEffectList(limit, offset)

// Individual resource group example returning second page of 20 results
superContestEffectList, err := contestsGroup.GetSuperContestEffectList(limit, offset)
```

##### Get URL for Super Contest Effect resource:
```go
// Main client example
superContestEffectURL := client.Contests.GetSuperContestEffectURL()

// Individual resource group example
superContestEffectURL := contestsGroup.GetSuperContestEffectURL()
```
