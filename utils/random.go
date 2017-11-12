package utils

import "github.com/google/uuid"

func GenerateIID() string {
	return uuid.New().String()
}
