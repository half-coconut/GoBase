package repository

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository/dao"
	"context"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}
func (r *UserRepository) FindByEmail(c context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	// 在这里操作缓存
}

func (r *UserRepository) Update(c context.Context, Id int64, nick_name, birthday, personal_profile string) (domain.User, error) {
	u, err := r.dao.Update(c, Id, nick_name, birthday, personal_profile)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		NickName:        u.NickName,
		Birthday:        u.Birthday,
		PersonalProfile: u.PersonalProfile,
	}, nil
}

func (r *UserRepository) FindById(c context.Context, id int64) (domain.User, error) {
	// 先从 cache 里面找
	// 再从 dao 里面找
	// 找到了回写 cache

	u, err := r.dao.FindById(c, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		NickName:        u.NickName,
		Birthday:        u.Birthday,
		PersonalProfile: u.PersonalProfile,
	}, nil
}
