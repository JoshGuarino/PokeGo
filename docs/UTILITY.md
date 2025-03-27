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
// Main client example returning first page of 20 results
languageList, err := client.Utility.GetLanguageList(20, 0)

// Individual resource group example returning second page of 20 results
languageList, err := utilityGroup.GetLanguageList(20, 20)
```
