# [Machines Group](https://pokeapi.co/docs/v2#machines-section)

The following are the resources available in the Machines group of the [PokeAPI](https://pokeapi.co/).

## [Machine](https://pokeapi.co/docs/v2#machine)

##### Get single Machine resource by ID:

```go
// Main client example with ID
machine, err := client.Machines.GetMachine("1")

// Individual resource group example with ID
machine, err := machinesGroup.GetMachine("1")
```

##### Get list of Machine resource:

```go
// Main client example 
machineList, err := client.Machines.GetMachineList(limit, offset)

// Individual resource group example 
machineList, err := machinesGroup.GetMachineList(limit, offset)
```

##### Get URL for Machine resource:

```go
// Main client example
machineURL := client.Machines.GetMachineURL()

// Individual resource group example
machineURL := machinesGroup.GetMachineURL()
```
