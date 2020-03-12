package util

import (
	"crypto/sha1"
)

// HashKey used to hash the key
func HashKey(k string) ([]byte, error) {

	h := sha1.New()

	if _, err := h.Write([]byte(k)); err != nil {
		return nil, err
	}
	val := h.Sum(nil)
	return val, nil
}
