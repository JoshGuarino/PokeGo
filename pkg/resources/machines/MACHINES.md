# [Machines Group](https://pokeapi.co/docs/v2#machines-section)
The following are the resources available in the Machines group of the [PokeAPI](https://pokeapi.co/).

## [Machine](https://pokeapi.co/docs/v2#machine)

##### Get single Machine resource by ID
```go
// Main client example with ID
machine, err := client.Machines.GetMachine("1")

// Individual resource group example with ID 
machine, err := machinesGroup.GetMachine("1")
```

##### Get list of Machine resource 
```go
// Main client example with no pagination options
machineList, err := client.Machines.GetMachineList(models.PaginationOptions{})

// Individual resource group example with pagination options
machineList, err := machinesGroup.GetMachineList(models.PaginationOptions{Limit: 20, Offset: 20})
```
