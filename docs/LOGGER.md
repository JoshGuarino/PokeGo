# Logger

The logger is initialized when the client is created or when a resource group is created.
Only one instance of the logger is created and shared between the client and all resource groups upon there initialization.
Any subsequent intializations will not create a new logger but reference the existing logger.
A pointer reference to the initialized logger is stored in the client and each resource group.
See singleton pattern [here](https://en.wikipedia.org/wiki/Singleton_pattern).
By default the logger is inactive and will not log any messages. To activate the logger, call the `SetActive()` method.


## Log Levels

The logger supports the following log levels:

- `Info`
- `Warn`
- `Error`
- `Debug`

It is set to `Info` by default.

## The following is a set of available methods for the logger used in the PokeGo client.

##### The info method logs a message at the `info` level.

```go
// Main client example
client.Log.Info("This is an info message")

// Individual resource group example
resourceGroup.Log.Info("This is an info message")
```

##### The warn method logs a message at the `warn` level.

```go
// Main client example
client.Log.Warn("This is a warning message")

// Individual resource group example
resourceGroup.Log.Warn("This is a warning message")
```

##### The error method logs a message at the `error` level.

```go
// Main client example
client.Log.Error("This is an error message")

// Individual resource group example
resourceGroup.Log.Error("This is an error message")
```

##### The debug method logs a message at the `debug` level.

```go
// Main client example
client.Log.Debug("This is a debug message")

// Individual resource group example
resourceGroup.Log.Debug("This is a debug message")
```

##### The messages method returns a slice of all logged messages.
```go
// Main client example
client.Log.Messages()


// Individual resource group example
resourceGroup.Log.Messages()
```

##### The logger can be cleared by calling the `Clear()` method:

```go
// Main client example
client.Log.Clear()

// Individual resource group example
resourceGroup.Log.Clear()
```

##### The active status of the logger can be set by calling the `setActive()` method:

```go
// Main client example
client.Log.SetActive(false)

// Individual resource group example
resourceGroup.Log.SetActive(false)
```

##### The active status of the logger can be checked by calling the `GetActive()` method:

```go
// Main client example
client.Logger.GetActive()

// Individual resource group example
resourceGroup.Logger.GetActive()
```

##### The level of the logger can be set by calling the `SetLevel()` method:

```go
// Main client example
client.Log.SetLevel(log.InfoLevel)

// Individual resource group example
resourceGroup.Log.SetLevel(log.InfoLevel)
```

##### The level of the logger can be checked by calling the `Level()` method:

```go
// Main client example
client.Log.Level()

// Individual resource group example
resourceGroup.Log.Level()
```
