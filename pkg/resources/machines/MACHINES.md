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
// Main client example returning first page of 20 results
machineList, err := client.Machines.GetMachineList(20, 0)

// Individual resource group example returning second page of 20 results
machineList, err := machinesGroup.GetMachineList(20, 20)
```
