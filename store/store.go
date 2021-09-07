package store

import (
	"fmt"
	"keyvaluestore/service"
	"os"
	"sync"
	"time"
)

type Store struct {
	cache map[string]string
	mut   *sync.RWMutex
}

// Create NewInMemoryStore and return new in memory store
func NewInMemoryStore() service.InMemoryService {
	return &Store{
		cache: make(map[string]string),
		mut:   &sync.RWMutex{},
	}
}

// GetValue returns key value pair from memory
func (s *Store) GetValue(req service.GetValueRequest) service.GetValueResponse {
	s.mut.RLock()
	defer s.mut.RUnlock()
	k := req.Key
	v, found := s.cache[k]
	if !found {
		return service.GetValueResponse{
			Error: fmt.Sprintf("value is not found for key: %s", k),
		}
	}
	return service.GetValueResponse{
		Key:   k,
		Value: v,
	}
}

//SetValue sets key value pair to memory
func (s *Store) SetValue(req service.SetValueRequest) service.SetValueResponse {
	k := req.Key
	v := req.Value
	_, found := s.cache[k]
	if !found {
		s.mut.Lock()
		s.cache[k] = v
		s.mut.Unlock()
		return service.SetValueResponse{
			Key:   k,
			Value: v,
		}
	}

	return service.SetValueResponse{
		Error: fmt.Sprintf("%s named key already is in memory", k),
	}

}

// DeleteValue deletes key value pair in memory
func (s *Store) DeleteValue(req service.GetValueRequest) service.DeleteValueResponse {
	k := req.Key
	_, found := s.cache[k]
	if found {
		s.mut.Lock()
		delete(s.cache, k)
		s.mut.Unlock()
		return service.DeleteValueResponse{
			Error: fmt.Sprintf("%s named key deleted from memory", k),
		}
	}
	return service.DeleteValueResponse{
		Error: fmt.Sprintf("%s named key is not in memory", k),
	}
}

// GetAll gets all key value pair in memory
func (s *Store) GetAll() service.GetAllResponse {
	kvs := make([]map[string]string, 0)
	kv := make(map[string]string)
	s.mut.RLock()
	defer s.mut.RUnlock()
	for key, value := range s.cache {
		kv[key] = value
	}
	kvs = append(kvs, kv)

	return service.GetAllResponse{
		Stores: kvs,
	}
}

// Writes data in memory to .txt file per 10 seconds
func (s *Store) WriteToFile() {

	ticker := time.NewTicker(10 * time.Second)

	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				file, err := os.OpenFile("keyvalue-db.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
				if err != nil {
					panic(err)
				}
				for key, value := range s.cache {
					file.WriteString("key: " + key + " value: " + value + "\n")
				}

				file.WriteString("----------------KEY VALUE ENDED---------------\n")
				file.Close()
			}
		}
	}()
}
