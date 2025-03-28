# Cache

The cache is initialized when the client is created or when a resource group is created.
Only one instance of the cache is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new cache but reference the existing cache.
A pointer reference to the initialized cache is stored in the client and each resource group.

## The following is a set of available methods for the cache used in the PokeGo client.

##### The cache can be cleared by calling the `Clear()` method:

```go
// Main client example
client.Cache.Clear()

// Individual resource group example
resourceGroup.Cache.Clear()
```

##### The active status of the cache can be set by calling the `setActive()` method:

```go
// Main client example
client.Cache.SetActive(false)

// Individual resource group example
resourceGroup.Cache.SetActive(false)
```

##### The active status of the cache can be checked by calling the `GetActive()` method:

```go
// Main client example
client.Cache.GetActive()

// Individual resource group example
resourceGroup.Cache.GetActive()
```

##### The expiration time of the cache can be set by calling the `SetExpiration()` method:

```go
// Main client example
client.Cache.SetExpiration(48 * time.Hour)

// Individual resource group example
resourceGroup.Cache.SetExpiration(48 * time.Hour)
```

##### The expiration time of the cache can be checked by calling the `GetExpiration()` method:

```go
// Main client example
client.Cache.GetExpiration()

// Individual resource group example
resourceGroup.Cache.GetExpiration()
```
