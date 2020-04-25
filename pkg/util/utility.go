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

// XorDistance will return the distance between two
// identifiers just like specified in the Kademlia protocol
func Xor(x []byte, y []byte) (ret [IDLENGTH]byte){
	for i := 0; i < IDLENGTH; i ++{
		ret[i] = x[i] ^ y[i]
	}
	return
}


func Equals(x []byte, y []byte) bool{
	for i := 0; i < IDLENGTH; i++ {
		if x[i] != y[i]{
			return false
		}
	}
	return true
}


func Less(x []byte, y []byte) bool{
	for i := 0; i < IDLENGTH; i++ {
		if x[i] != y[i] {
			return x[i] < y[i]
		}
	}
	return false
}

func PrefixLen(id []byte) (ret int) {
	for i := 0; i < IDLENGTH; i++ {
		for j := 0; j < 8; j++ {
			if (id[i] >> uint8(7 - j)) & 0x1 != 0 {
				return i * 8 + j
			}
		}
	}
	return IDLENGTH * 8 - 1
}