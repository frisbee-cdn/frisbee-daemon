package common

import (
	"bytes"
	"crypto/sha1"
	"math/bits"
)

// NodeID binary number which represents the hash of the random node id
type NodeID []byte

// HashKey used to hash the key
func HashKey(id ID) (NodeID, error) {

	h := sha1.New()

	if _, err := h.Write([]byte(id)); err != nil {
		return nil, err
	}
	val := h.Sum(nil)
	return val, nil
}

// XorDistance will return the distance between two
// identifiers just like specified in the Kademlia protocol
func XOR(x []byte, y []byte) []byte {
	c := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		c[i] = x[i] ^ y[i]
	}
	return c
}

// Equals returns whether keys are equal in this key space
func (id NodeID) Equals(other NodeID) bool {
	return bytes.Equal(id, other)
}

// Less returns whether the first key is smaller than the second one
func (id NodeID) Less(other NodeID) bool {
	return bytes.Compare(id, other) < 0
}

// ZeroPrefixLen is used to find the numbe of leading 0's
func ZeroPrefixLen(id []byte) (ret int) {
	for i, b := range id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(id) * 8
}

func CommonPrefixLen(x, y []byte) int {
	return ZeroPrefixLen(XOR(x, y))
}
