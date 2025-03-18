# [Encounters Group](https://pokeapi.co/docs/v2#encounters-section)
The following are the resources available in the Encounters group of the [PokeAPI](https://pokeapi.co/).

## [EncounterMethod](https://pokeapi.co/docs/v2#encounter-methods)

##### Get single Encounter Method resource by name or ID
```go
// Main client example with ID
encounterMethod, err := client.Encounters.GetEncounterMethod("1")

// Individual resource group example with Name
encounterMethod, err := encountersGroup.GetEncounterMethod("walk")
```

##### Get list of Encounter Method resource
```go
// Main client example with no pagination options
encounterMethodList, err := client.Encounters.GetEncounterMethodList(models.PaginationOptions{})

// Individual resource group example with pagination options
encounterMethodList, err := encountersGroup.GetEncounterMethodList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [EncounterCondition](https://pokeapi.co/docs/v2#encounter-conditions)

##### Get single Encounter Condition resource by name or ID
```go
// Main client example with ID
encounterCondition, err := client.Encounters.GetEncounterCondition("1")

// Individual resource group example with Name
encounterCondition, err := encountersGroup.GetEncounterCondition("swarm")
```

##### Get list of Encounter Condition resource
```go
// Main client example with no pagination options
encounterConditionList, err := client.Encounters.GetEncounterConditionList(models.PaginationOptions{})

// Individual resource group example with pagination options
encounterConditionList, err := encountersGroup.GetEncounterConditionList(models.PaginationOptions{Limit: 20, Offset: 20})
```

## [EncounterConditionValue](https://pokeapi.co/docs/v2#encounter-condition-values)

##### Get single Encounter Condition Value resource by name or ID
```go
// Main client example with ID
encounterConditionValue, err := client.Encounters.GetEncounterConditionValue("1")

// Individual resource group example with Name
encounterConditionValue, err := encountersGroup.GetEncounterConditionValue("swarm-yes")
```

##### Get list of Encounter Condition Value resource
```go
// Main client example with no pagination options
encounterConditionValueList, err := client.Encounters.GetEncounterConditionValueList(models.PaginationOptions{})

// Individual resource group example with pagination options
encounterConditionValueList, err := encountersGroup.GetEncounterConditionValueList(models.PaginationOptions{Limit: 20, Offset: 20})
```
