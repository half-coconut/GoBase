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
	FindByWechat(c context.Context, openID string) (domain.User, error)
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
	if err == nil {
		// 有数据，返回 u
		return u, nil
	}

	ue, err := r.dao.FindById(c, id)
	if err != nil {
		return domain.User{}, err
	}
	u = r.entityToDomain(ue)
	// 同步调用 set，不使用 go routine 方式,
	//_ = r.cache.Set(c, u)

	// 异步调用 set
	go func() {
		_ = r.cache.Set(c, u)
		//err = r.cache.Set(c, u)
		//if err != nil {
		//	// 怎么处理
		//}
	}()
	return u, nil
}

func (r *CachedUserRepository) FindByWechat(c context.Context, openID string) (domain.User, error) {
	u, err := r.dao.FindByWechat(c, openID)
	return r.entityToDomain(u), err
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
		Phone:           u.Phone.String,
		NickName:        u.NickName.String,
		Birthday:        birthday,                 // 前端输入 1990-01-01 需要转化吗？
		PersonalProfile: u.PersonalProfile.String, // 200个字符
		WechatInfo: domain.WechatInfo{
			UnionID: u.WechatUnionID.String,
			OpenID:  u.WechatOpenID.String,
		},
		Ctime: time.UnixMilli(u.Ctime),
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
		WechatOpenID: sql.NullString{
			String: u.WechatInfo.OpenID,
			Valid:  u.WechatInfo.OpenID != "",
		},
		WechatUnionID: sql.NullString{
			String: u.WechatInfo.UnionID,
			Valid:  u.WechatInfo.UnionID != "",
		},
	}
}
