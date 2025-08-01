# [Encounters Group](https://pokeapi.co/docs/v2#encounters-section)

The following are the resources available in the Encounters group of the [PokeAPI](https://pokeapi.co/).

## [EncounterMethod](https://pokeapi.co/docs/v2#encounter-methods)

##### Get single Encounter Method resource by name or ID:

```go
// Main client example with ID
encounterMethod, err := client.Encounters.GetEncounterMethod("1")

// Individual resource group example with Name
encounterMethod, err := encountersGroup.GetEncounterMethod("walk")
```

##### Get list of Encounter Method resource:

```go
// Main client example 
encounterMethodList, err := client.Encounters.GetEncounterMethodList(limit, offset)

// Individual resource group example 
encounterMethodList, err := encountersGroup.GetEncounterMethodList(limit, offset)
```

##### Get URL for Encounter Method resource:
```go
// Main client example
encounterMethodURL := client.Encounters.GetEncounterMethodURL()

// Individual resource group example
encounterMethodURL := encountersGroup.GetEncounterMethodURL()
```

## [EncounterCondition](https://pokeapi.co/docs/v2#encounter-conditions)

##### Get single Encounter Condition resource by name or ID:

```go
// Main client example with ID
encounterCondition, err := client.Encounters.GetEncounterCondition("1")

// Individual resource group example with Name
encounterCondition, err := encountersGroup.GetEncounterCondition("swarm")
```

##### Get list of Encounter Condition resource:

```go
// Main client example 
encounterConditionList, err := client.Encounters.GetEncounterConditionList(limit, offset)

// Individual resource group example
encounterConditionList, err := encountersGroup.GetEncounterConditionList(limit, offset)
```

##### Get URL for Encounter Condition resource:
```go
// Main client example
encounterConditionURL := client.Encounters.GetEncounterConditionURL()

// Individual resource group example
encounterConditionURL := encountersGroup.GetEncounterConditionURL()
```

## [EncounterConditionValue](https://pokeapi.co/docs/v2#encounter-condition-values)

##### Get single Encounter Condition Value resource by name or ID:

```go
// Main client example with ID
encounterConditionValue, err := client.Encounters.GetEncounterConditionValue("1")

// Individual resource group example with Name
encounterConditionValue, err := encountersGroup.GetEncounterConditionValue("swarm-yes")
```

##### Get list of Encounter Condition Value resource:

```go
// Main client example 
encounterConditionValueList, err := client.Encounters.GetEncounterConditionValueList(limit, offset)

// Individual resource group example 
encounterConditionValueList, err := encountersGroup.GetEncounterConditionValueList(limit, offset)
```

##### Get URL for Encounter Condition Value resource:
```go
// Main client example
encounterConditionValueURL := client.Encounters.GetEncounterConditionValueURL()

// Individual resource group example
encounterConditionValueURL := encountersGroup.GetEncounterConditionValueURL()
```
