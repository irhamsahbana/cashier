package validator

import (
	"github.com/google/uuid"
)

func IsUUID(u string) error {
    _, err := uuid.Parse(u)
    return err
 }