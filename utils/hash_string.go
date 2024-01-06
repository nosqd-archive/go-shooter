package utils

import "hash/fnv"

func HashStringToUint32(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		// Handle error
	}
	return h.Sum32()
}
