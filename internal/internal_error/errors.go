package internalerror

import (
	"errors"

	"gorm.io/gorm"
)

var RepositoryErr error = errors.New("error in repository")

func ProccessError(err error) error {
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("internal server error")
		}
		return errors.New("record not found")
	}
	return nil
}