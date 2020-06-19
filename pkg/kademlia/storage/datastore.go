package storage

import (
	"errors"
	"sync"
)

// ErrNotFound marks a not found error
var ErrNotFound = errors.New("Not Found")

//MapDatastore implements a simple store with the web resouces kept in memory
type MapDatastore struct {
	Storage

	values map[string][]byte
	mutex  *sync.RWMutex
}

// NewMapDatastore creates a MemoryStorage prepared
func NewMapDatastore() *MapDatastore {
	return &MapDatastore{
		values: make(map[string][]byte),
	}
}

// Get retrieves the web resource using its uri address
func (m *MapDatastore) Get(key string) ([]byte, error) {

	val, err := m.values[key]

	if !err {
		return nil, ErrNotFound
	}
	return val, nil
}

// Put saves the passwed web resource in the memory store
func (m *MapDatastore) Put(key string, content []byte) error {
	m.values[key] = content
	return nil
}

// Delete removes the web resource along with its uri address from the memory store
func (m *MapDatastore) Delete(key string) error {
	delete(m.values, key)
	return nil
}

// Contains checks the existence of an particular data in the storage
func (m *MapDatastore) Contains(key string) bool {
	_, found := m.values[key]
	return found
}
