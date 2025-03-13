# PokeGo

PokeGo is an Golang API wrapper for the PokéAPI.

## Todo
- Add documentation
- Improve testing

## Setup
There are 2 options for using PokeGo. You can either use the main client or create individual resource groups seperately. 
The main client will initialize all resource groups for you. If you choose to use individual resource groups, 
you will need to initialize each group separately.

Main client setup:
```go 
import (
    pokego "github.com/JoshGuarino/PokeGo/pkg"
)

client := pokego.NewClient()
```
Individual resource group setup:
```go
import (
    "github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

pokemonGroup := pokemon.NewPokemonGroup()
```


## Caching
PokeGo uses a simple in-memory cache to store API responses. This is to reduce the number of requests made to the PokeAPI. 
The cache is set to expire after 24 hours as resources in the PokeAPI are mostly static. 
The cache is initialized when the client is created or when a resource group is created
Only one instance of the cache is created and shared between client and all resource groups upon initialization.
Any subsequent intializations will not create a new cache but reference the existing cache.
A pointer reference to the initialized cache is stored in the client and each resource group.

### Clearing cache
The cache can be cleared by calling the `Clear()` method on the cache. 
```go
// Main client example
client.Cache.Clear()

// Individual resource group example
resourceGroup.Cache.Clear()
```

### Disabling Cache
The active status of the cache can be set by calling the `setActive()` method on the cache. 
```go
// Main client example
client.Cache.SetActive(false)

// Individual resource group example
resourceGroup.Cache.SetActive(false)
```
### Custom expiration time
The expiration time of the cache can be set by calling the `SetExpiration()` method on the cache. 
```go
// Main client example
client.Cache.SetExpiration(48 * time.Hour)

// Individual resource group example
resourceGroup.Cache.SetExpiration(48 * time.Hour)
```

## License
BSD 3-Clause License

Copyright (c) 2025, Josh Guarino

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
