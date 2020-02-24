package storage

import (
	"sync"
)

//MemoryStorage implements a simple store with the web resouces kept in memory
type MemoryStorage struct {
	Storage

	resources map[string]WebResources
	mux       *sync.RWMutex
}

//Get retrieves the web resource using its uri address
func (m *MemoryStorage) Get(key string) ([]byte, error) {

	return nil, nil
}

//Save saves the passwed web resource in the memory store
func (m *MemoryStorage) Save(key string, content []byte) error {

	return nil
}

//Delete removes the web resource along with its uri address from the memory store
func (m *MemoryStorage) Delete(key string) error {

	return nil
}
