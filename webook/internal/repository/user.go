package repository

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository/cache"
	"GoBase/webook/internal/repository/dao"
	"context"
	"database/sql"
	"time"
)

var (
	ErrUserDuplicate = dao.ErrUserDuplicate
	ErrUserNotFound  = dao.ErrUserNotFound
)

type UserRepository interface {
	Create(c context.Context, u domain.User) error
	// Update 更新数据，只有非 0 值才会更新
	Update(c context.Context, u domain.User) error
	FindByEmail(c context.Context, email string) (domain.User, error)
	FindByPhone(c context.Context, phone string) (domain.User, error)
	FindById(c context.Context, id int64) (domain.User, error)
}

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewUserRepository(d dao.UserDAO, c cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   d,
		cache: c,
	}
}

func (r *CachedUserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, dao.User{
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Password: u.Password,
	})
	// 在这里操作缓存
}

func (r *CachedUserRepository) Update(c context.Context, u domain.User) error {
	err := r.dao.UpdateNonZeroFields(c, r.domainToEntity(u))
	if err != nil {
		return err
	}
	return r.cache.Delete(c, u.Id)
}
func (r *CachedUserRepository) FindByEmail(c context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(c, email)
	return r.entityToDomain(u), err
}

func (r *CachedUserRepository) FindByPhone(c context.Context, phone string) (domain.User, error) {
	u, err := r.dao.FindByPhone(c, phone)
	return r.entityToDomain(u), err
}

func (r *CachedUserRepository) FindById(c context.Context, id int64) (domain.User, error) {
	// 先从 cache 里面找
	// 再从 dao 里面找
	// 找到了回写 cache
	u, err := r.cache.Get(c, id)
	if err != nil {
		return u, err
	}

	ue, err := r.dao.FindById(c, id)
	if err != nil {
		return domain.User{}, err
	}
	u = r.entityToDomain(ue)
	go func() {
		err = r.cache.Set(c, u)
		if err != nil {
			// 怎么处理
		}
	}()
	return u, nil
}

func (r *CachedUserRepository) entityToDomain(u dao.User) domain.User {
	var birthday time.Time
	if u.Birthday.Valid {
		birthday = time.UnixMilli(u.Birthday.Int64)
	}
	return domain.User{
		Id:              u.Id,
		Email:           u.Email.String,
		Password:        u.Password,
		NickName:        u.NickName.String,
		Birthday:        birthday,                 // 前端输入 1990-01-01 需要转化吗？
		PersonalProfile: u.PersonalProfile.String, // 200个字符
		Ctime:           time.UnixMilli(u.Ctime),
	}
}

func (r *CachedUserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Birthday: sql.NullInt64{
			Int64: u.Birthday.UnixMilli(),
			Valid: !u.Birthday.IsZero(),
		},
		NickName: sql.NullString{
			String: u.NickName,
			Valid:  u.NickName != "",
		},
		PersonalProfile: sql.NullString{
			String: u.PersonalProfile,
			Valid:  u.PersonalProfile != "",
		},
		Password: u.Password,
	}
}
