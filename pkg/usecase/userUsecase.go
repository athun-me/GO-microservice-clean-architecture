package usecase

import (
	"athun/pkg/domain"
	interfaces "athun/pkg/repository/interface"
	userCase "athun/pkg/usecase/interface"
)

type userUseCase struct {
	Repo interfaces.UserRepo
}

func NewUserUseCase(repo interfaces.UserRepo) userCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}

func (use *userUseCase) Register(user domain.User) error {
	err := use.Repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
