# Environment

The environment is initialized when the client is created or when a resource group is created.
Only one instance of the environment is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new environment but reference the existing environment.
A pointer reference to the initialized environment is stored in the client and each resource group.
See singleton pattern [here](https://en.wikipedia.org/wiki/Singleton_pattern).

## Switching Environments

You can easily switch between the `production` and `staging` environments.
The main client and individual resource groups all have a environment struct pounter 
that can be used to access environment methods.

## The following is a set of available methods for the environment used in the PokeGo client.

##### The domain of the environment can be checked by calling the `Domain()` method:

```go
// Main client example
client.Env.Domain()

// Individual resource group example
resourceGroup.Env.Domain()
```

##### The base URL of the environment can be checked by calling the `URL()` method:

```go
// Main client example
client.Env.URL()

// Individual resource group example
resourceGroup.Env.URL()
```

##### The environment can be switched to the `production` environment by calling the `SetProd()` method:

```go
// Main client example
client.Env.SetProd()

// Individual resource group example
resourceGroup.Env.SetProd()
```

##### The environment can be switched to the `staging` environment by calling the `SetStage()` method:

```go
// Main client example
client.Env.SetStage()

// Individual resource group example
resourceGroup.Env.SetStage()
```
