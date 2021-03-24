package model

import (
	"crypto/sha256"
	"fmt"
)

// 64 length hex string
type ID string

// CreateID for entity type & identifier
func CreateID(entityType string, uriOrIDOrTitle interface{}) ID {
	hash := sha256.New()
	hash.Write([]byte(entityType))
	hash.Write([]byte(fmt.Sprintf("%v", uriOrIDOrTitle)))
	return ID(fmt.Sprintf("%x", hash.Sum(nil)))
}
