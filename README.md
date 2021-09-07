# Key-Value Store
In memory key value store application with API service

# Installation
```bash
go get -u github.com/batuhankecici/keyvaluestore
```

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

Get handler response key value pair from memory
```zsh
localhost:****/get
```
Set handler sets key value pair to memory
```zsh
localhost:****/set
```
Delete handler deletes key value pair to memory
```zsh
localhost:****/delete
```
Getall handler gets all key value pair in memory
```zsh
localhost:****/getall
```

# Benchmark Tests
**Test Codes:**

```go
func BenchmarkCreateHttpHandler(b *testing.B) {
	ims := inMemoryStore.NewInMemoryStore()

	for i := 0; i < b.N; i++ {
		CreateHTTPHandler(ims)
	}
}

func BenchmarkSet(b *testing.B) {
	ims := NewInMemoryStore()

	k := "batuhan"
	v := "kecici"

	setValueReq := service.SetValueRequest{
		Key:   k,
		Value: v,
	}

	for i := 0; i < b.N; i++ {
		ims.SetValue(setValueReq)
	}
}

func BenchmarkGet(b *testing.B) {
	ims := NewInMemoryStore()

	k := "batuhan"
	v := "kecici"

	setValueReq := service.SetValueRequest{
		Key:   k,
		Value: v,
	}
	getValueReq := service.GetValueRequest{
		Key: k,
	}
	ims.SetValue(setValueReq)
	for i := 0; i < b.N; i++ {
		ims.GetValue(getValueReq)
	}
}

```

**Results:**
|Function|Time|Bytes Allocated|Objects Allocated|
|:-------|:--:|:-------------:|:---------------:|
|CreateHTTPHandler|674.7 ns/op|624 B/op|7 allocs/op|
|Set|202.1 ns/op|64 B/op|2 allocs/op|
|Get|28.20 ns/op|0 B/op|0 allocs/op|