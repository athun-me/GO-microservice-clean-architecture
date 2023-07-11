package interfaces

import "athun/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) error
}
