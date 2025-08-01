# [Utility Group](https://pokeapi.co/docs/v2#utility-section)

The following are the utility functions and features of the PokeGo API client.

## [Language](https://pokeapi.co/docs/v2#languages)

##### Get single Language resource by Name or ID:

```go
// Main client example with ID
language, err := client.Utility.GetLanguage("1")

// Individual resource group example with Name
language, err := utilityGroup.GetLanguage("en")
```

##### Get list of Language resource:

```go
// Main client example 
languageList, err := client.Utility.GetLanguageList(limit, offset)

// Individual resource group example 
languageList, err := utilityGroup.GetLanguageList(limit, offset)
```

##### Get URL for Language resource:

```go
// Main client example
languageURL := client.Utility.GetLanguageURL()

// Individual resource group example
languageURL := utilityGroup.GetLanguageURL()
```
