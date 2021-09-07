# Key-Value Store
In memory key value store application with API service

# Installation
```bash
go get -u github.com/batuhankecici/Golang/keyvaluestore
```

# Usage

```go
package main

import (
	kvstore "github.com/batuhankecici/Golang/keyvaluestore"
)

func main() {
	// create new in memory store
	ims := kvstore.store.NewInMemoryStore()
	// create http handler
	h := kvstore.transport.CreateHTTPHandler(ims)
}
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