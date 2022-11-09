package validation

import (
	"fmt"

	"vingolds.ch/errgo/validation_errors"
)

func RowNotFound() error {
	return fmt.Errorf("unexpected: %w", validation_errors.SqlNoSuchRow)
}

func MemberNotFound(memberId int, err error) error {
	return validation_errors.MkMemberNotFound(fmt.Sprintf("could not load member id=%v", memberId), err)
}
