package usecases

import "cleango/app/domains"

type UserRepository interface {
	Save(domains.User) (string, error)
	FindAll() (domains.Users, error)
	// DeleteByID(string) error
	// FindByID(string) (domains.User, error)
	// UpdateByID(string, domains.User) error
}
