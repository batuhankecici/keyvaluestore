# Key-Value Store
In memory key value store application with API service

# Installation
```bash
go get -u github.com/batuhankecici/keyvaluestore
```

# Dependencies

Http Server adress using :8080 port.You should send a request to "localhost:8080"
# Usage

```go
package main

import (
	kvstore "github.com/batuhankecici/keyvaluestore"
)

func main() {
	// create new in memory store
	ims := kvstore.store.NewInMemoryStore()
	// create http handler
	h := kvstore.transport.CreateHTTPHandler(ims)
}
```
# Endpoints

Keyvaluestore has 4 endpoints.
Get handler response key value pair from memory
```zsh
	localhost:8080/get
```
Set handler sets key value pair to memory
```zsh
	localhost:8080/set
```
Delete handler deletes key value pair to memory
```zsh
	localhost:8080/delete
```
Getall handler gets all key value pair in memory
```zsh
	localhost:8080/getall
```

# Benchmark Tests
**Test Code:**

```go
func BenchmarkCreateHttpHandler(b *testing.B) {
	ims := inMemoryStore.NewInMemoryStore()

	for i := 0; i < b.N; i++ {
		CreateHTTPHandler(ims)
	}
}

```

**Results:**
|Function|Time|Bytes Allocated|Objects Allocated|
|:-------|:--:|:-------------:|:---------------:|
|CreateHTTPHandler|674.7 ns/op|624 B/op|7 allocs/op|