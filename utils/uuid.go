package utils

import uuid "github.com/satori/go.uuid"

func GeneratorUUID() string {
	id := uuid.NewV4()
	return id.String()
}
