package interfaces

import "athun/pkg/domain"

type UserRepo interface {
	Create(user domain.User) error
}
