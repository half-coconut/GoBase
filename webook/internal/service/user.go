package service

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(c context.Context, u domain.User) error {

	return svc.repo.Create(c, u)
}
