package utils

import "fmt"

type Hash uint64

func HashString(key string) Hash {
	var hash Hash = 5381

	for c := range key {
		hash = (Hash(hash << 5) + hash) + Hash(c)
	}

	return hash
}

func (h Hash) String() string {
	return "0x" + fmt.Sprintf("%X", uint8(h))
}