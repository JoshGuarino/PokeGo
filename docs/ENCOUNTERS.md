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
// Main client example returning first page of 20 results
encounterMethodList, err := client.Encounters.GetEncounterMethodList(20, 0)

// Individual resource group example returning second page of 20 results
encounterMethodList, err := encountersGroup.GetEncounterMethodList(20, 20)
```

##### Get URL for Encounter Method resource:
```go
// Main client example
encounterMethodURL, err := client.Encounters.GetEncounterMethodURL()

// Individual resource group example
encounterMethodURL, err := encountersGroup.GetEncounterMethodURL()
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
// Main client example returning first page of 20 results
encounterConditionList, err := client.Encounters.GetEncounterConditionList(20, 0)

// Individual resource group example returning second page of 20 results
encounterConditionList, err := encountersGroup.GetEncounterConditionList(20, 20)
```

##### Get URL for Encounter Condition resource:
```go
// Main client example
encounterConditionURL, err := client.Encounters.GetEncounterConditionURL()

// Individual resource group example
encounterConditionURL, err := encountersGroup.GetEncounterConditionURL()
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
// Main client example returning first page of 20 results
encounterConditionValueList, err := client.Encounters.GetEncounterConditionValueList(20, 0)

// Individual resource group example returning second page of 20 results
encounterConditionValueList, err := encountersGroup.GetEncounterConditionValueList(20, 20)
```

##### Get URL for Encounter Condition Value resource:
```go
// Main client example
encounterConditionValueURL, err := client.Encounters.GetEncounterConditionValueURL()

// Individual resource group example
encounterConditionValueURL, err := encountersGroup.GetEncounterConditionValueURL()
```
