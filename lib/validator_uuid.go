package lib

import (
	"github.com/google/uuid"
)

func ValidatorUUID(u string) (bool, error) {
    _, err := uuid.Parse(u)
    return (err == nil), err
 }