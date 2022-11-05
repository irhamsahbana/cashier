package validator

import (
	"github.com/google/uuid"
)

func Uuid(u string) error {
	_, err := uuid.Parse(u)
	return err
}
