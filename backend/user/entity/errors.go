package entity

import "errors"

var (
	ErrUserNotExist         = errors.New("User does not exist")
	ErrOrganizationNotExist = errors.New("Organization does not exist")
)
