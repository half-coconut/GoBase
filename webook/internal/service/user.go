package service

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

/**
go install go.uber.org/mock/mockgen@latest
mockgen -source=webook/internal/service/user.go -package=svcmocks -destination=webook/internal/service/mocks/user.mock.go

需要面向 interface 编程，重新改造后，再进行 mock

*/

var ErrUserDuplicate = repository.ErrUserDuplicate
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不正确")

type UserService interface {
	Login(c context.Context, email, password string) (domain.User, error)
	SignUp(c context.Context, u domain.User) error
	FindOrCreate(c context.Context, phone string) (domain.User, error)
	Edit(c context.Context, u domain.User) error
	Profile(c context.Context, id int64) (domain.User, error)
	UpdateNonSensitiveInfo(c context.Context, u domain.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) Login(c context.Context, email, password string) (domain.User, error) {
	// 先找用户
	u, err := svc.repo.FindByEmail(c, email)
	// 如果用户不存在，没找到
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	// 这里处理系统错误
	if err != nil {
		return domain.User{}, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// 输入密码错误，这里可以 DEBUG 打日志
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *userService) SignUp(c context.Context, u domain.User) error {
	// 考虑将加密放在这里
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(c, u)
}
func (svc *userService) FindOrCreate(c context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindByPhone(c, phone)
	if err != repository.ErrUserNotFound {
		return u, err
	}
	err = svc.repo.Create(c, domain.User{Phone: phone})
	if err != nil && err != repository.ErrUserDuplicate {
		return domain.User{}, err
	}
	return svc.repo.FindByPhone(c, phone)
}

func (svc *userService) Edit(c context.Context, u domain.User) error {
	// 编辑 Profile
	return svc.repo.Update(c, u)

}
func (svc *userService) UpdateNonSensitiveInfo(c context.Context, u domain.User) error {
	u.Email = ""
	u.Phone = ""
	u.Password = ""
	return svc.repo.Update(c, u)
}

func (svc *userService) Profile(c context.Context, id int64) (domain.User, error) {
	// 查询 Profile
	u, err := svc.repo.FindById(c, id)
	if err != nil {
		return domain.User{}, err
	}
	return u, err
}
