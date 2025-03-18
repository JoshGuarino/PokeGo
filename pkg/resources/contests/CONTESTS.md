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
// Main client example with no pagination options
contestTypeList, err := client.Contests.GetContestTypeList(models.PaginationOptions{})

// Individual resource group example with pagination options
contestTypeList, err := contestsGroup.GetContestTypeList(models.PaginationOptions{Limit: 20, Offset: 20})
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
// Main client example with no pagination options
contestEffectList, err := client.Contests.GetContestEffectList(models.PaginationOptions{})

// Individual resource group example with pagination options
contestEffectList, err := contestsGroup.GetContestEffectList(models.PaginationOptions{Limit: 20, Offset: 20})
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
// Main client example with no pagination options
superContestEffectList, err := client.Contests.GetSuperContestEffectList(models.PaginationOptions{})

// Individual resource group example with pagination options
superContestEffectList, err := contestsGroup.GetSuperContestEffectList(models.PaginationOptions{Limit: 20, Offset: 20})
```
