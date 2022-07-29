package validator

import (
	"github.com/google/uuid"
)

func IsUUID(u string) (bool, error) {
    _, err := uuid.Parse(u)
    return (err == nil), err
 }