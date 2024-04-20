package id

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Generate(resourceType string) string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("firestore: crypto/rand.Read error: %v", err))
	}

	r := base64.RawURLEncoding.EncodeToString(b)
	return fmt.Sprintf("%s_%s", r, resourceType)
}
