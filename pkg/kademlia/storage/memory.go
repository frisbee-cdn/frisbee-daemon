package storage

import (
	"fmt"
	"sync"
)

//MemoryStorage implements a simple store with the web resouces kept in memory
type MemoryStorage struct {
	Storage

	Resources map[string]WebResource
	mutex     *sync.RWMutex
}

//NewMemoryStorage creates a MemoryStorage prepared
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Resources: make(map[string]WebResource),
	}
}

//Get retrieves the web resource using its uri address
func (m *MemoryStorage) Get(key string) ([]byte, error) {

	for uri, wb := range m.Resources {
		if key == uri {
			return wb.Content, nil
		}
	}
	return nil, nil
}

//Save saves the passwed web resource in the memory store
func (m *MemoryStorage) Save(key string, content []byte) error {

	return nil
}

//Delete removes the web resource along with its uri address from the memory store
func (m *MemoryStorage) Delete(key string) error {

	for uri := range m.Resources {
		if uri == key {
			delete(m.Resources, uri)
			return nil
		}
	}
	return fmt.Errorf("No web resource with the specified URI found")
}
