package storage

// Storage interface which defines tha available operations
// that we can use to manipulate the storage.
type Storage interface {
	Get(key string) ([]byte, error)
	Put(key string, content []byte) error
	Delete(key string) error
	Contains(key string) bool
}
