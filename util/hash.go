package util

import "hash/fnv"

var (
	hash = fnv.New32()
)

func Hash(s string) uint32 {
	hash.Reset()
	hash.Write([]byte(s))
	return hash.Sum32()
}
