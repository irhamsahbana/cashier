package validator

import (
	"github.com/google/uuid"
)

func IsUUID(u string) error {
	_, err := uuid.Parse(u)
	return err
}

func In[T comparable](slice []T, s T) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
