# [Moves Group](https://pokeapi.co/docs/v2#moves-section)
The following are the resources available in the Moves group of the [PokeAPI](https://pokeapi.co/).

## [Move](https://pokeapi.co/docs/v2#move)

##### Get single Move resource by Name or ID
```go
// Main client example with ID
move, err := client.Moves.GetMove("1")

// Individual resource group example with Name
move, err := movesGroup.GetMove("pound")
```

##### Get list of Move resource 
```go
// Main client example with no pagination options
moveList, err := client.Moves.GetMoveList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveList, err := movesGroup.GetMoveList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Ailment](https://pokeapi.co/docs/v2#move-ailments)

##### Get single Move Ailment resource by Name or ID
```go
// Main client example with ID
moveAilment, err := client.Moves.GetMoveAilment("1")

// Individual resource group example with Name
moveAilment, err := movesGroup.GetMoveAilment("paralysis")
```

##### Get list of Move Ailment resource 
```go
// Main client example with no pagination options
moveAilmentList, err := client.Moves.GetMoveAilmentList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveAilmentList, err := movesGroup.GetMoveAilmentList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Battle Style](https://pokeapi.co/docs/v2#move-battle-styles)

##### Get single Move Battle Style resource by Name or ID
```go
// Main client example with ID
moveBattleStyle, err := client.Moves.GetMoveBattleStyle("1")

// Individual resource group example with Name
moveBattleStyle, err := movesGroup.GetMoveBattleStyle("attack")
```

##### Get list of Move Battle Style resource 
```go
// Main client example with no pagination options
moveBattleStyleList, err := client.Moves.GetMoveBattleStyleList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveBattleStyleList, err := movesGroup.GetMoveBattleStyleList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Category](https://pokeapi.co/docs/v2#move-categories)

##### Get single Move Category resource by Name or ID
```go
// Main client example with ID
moveCategory, err := client.Moves.GetMoveCategory("1")

// Individual resource group example with Name
moveCategory, err := movesGroup.GetMoveCategory("damage")
```

##### Get list of Move Category resource 
```go
// Main client example with no pagination options
moveCategoryList, err := client.Moves.GetMoveCategoryList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveCategoryList, err := movesGroup.GetMoveCategoryList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Damage Class](https://pokeapi.co/docs/v2#move-damage-classes)

##### Get single Move Damage Class resource by Name or ID
```go
// Main client example with ID
moveDamageClass, err := client.Moves.GetMoveDamageClass("1")

// Individual resource group example with Name
moveDamageClass, err := movesGroup.GetMoveDamageClass("status")
```

##### Get list of Move Damage Class resource 
```go
// Main client example with no pagination options
moveDamageClassList, err := client.Moves.GetMoveDamageClassList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveDamageClassList, err := movesGroup.GetMoveDamageClassList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Learn Method](https://pokeapi.co/docs/v2#move-learn-methods)

##### Get single Move Learn Method resource by Name or ID
```go
// Main client example with ID
moveLearnMethod, err := client.Moves.GetMoveLearnMethod("1")

// Individual resource group example with Name
moveLearnMethod, err := movesGroup.GetMoveLearnMethod("level-up")
```

##### Get list of Move Learn Method resource 
```go
// Main client example with no pagination options
moveLearnMethodList, err := client.Moves.GetMoveLearnMethodList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveLearnMethodList, err := movesGroup.GetMoveLearnMethodList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [Move Target](https://pokeapi.co/docs/v2#move-targets)

##### Get single Move Target resource by Name or ID
```go
// Main client example with ID
moveTarget, err := client.Moves.GetMoveTarget("1")

// Individual resource group example with Name
moveTarget, err := movesGroup.GetMoveTarget("specific-move")
```

##### Get list of Move Target resource 
```go
// Main client example with no pagination options
moveTargetList, err := client.Moves.GetMoveTargetList(models.PaginationOptions{})

// Individual resource group example with pagination options
moveTargetList, err := movesGroup.GetMoveTargetList(models.PaginationOptions{Limit: 20, Offset: 20})
```
