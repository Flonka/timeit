// Package uuid provides simple string uuids
package uuid

import (
	"crypto/rand"
	"fmt"
)

// NewhexUUid creates a random uuid
func NewHexUuid(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
