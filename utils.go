package fsxml

import uuid "github.com/satori/go.uuid"

// String return a pointer to a string
func String(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// UUIDFromString returns the corresponding
// UUIDv4 base of a string representation
func UUIDFromString(s string) *uuid.UUID {
	uuid, _ := uuid.FromString(s)

	return &uuid
}
