package repository

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository/cache"
	cachemocks "GoBase/webook/internal/repository/cache/mocks"
	"GoBase/webook/internal/repository/dao"
	daomocks "GoBase/webook/internal/repository/dao/mocks"
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

func TestCachedUserRepository_Create(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		c context.Context
		u domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			if err := r.Create(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCachedUserRepository_FindByEmail(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		c     context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			got, err := r.FindByEmail(tt.args.c, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCachedUserRepository_FindById(t *testing.T) {
	now := time.Now()
	now = time.UnixMilli(now.UnixMilli())
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache)
		ctx      context.Context
		id       int64
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "缓存未命中，查询成功",
			mock: func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache) {
				// 缓存未命中，查了缓存，但是没结果
				c := cachemocks.NewMockUserCache(ctrl)
				c.EXPECT().Get(gomock.Any(), int64(123)).
					Return(domain.User{}, cache.ErrKeyNotExist)

				d := daomocks.NewMockUserDAO(ctrl)
				d.EXPECT().FindById(gomock.Any(), int64(123)).
					Return(dao.User{
						Id: 123,
						Email: sql.NullString{
							String: "123@qq.com",
							Valid:  true,
						},
						Password: "this is password",
						Phone: sql.NullString{
							String: "12300001234",
							Valid:  true,
						},
						Ctime: now.UnixMilli(), // 使用毫秒数
						Utime: now.UnixMilli(),
					}, nil)
				c.EXPECT().Set(gomock.Any(), domain.User{
					Id:       123,
					Email:    "123@qq.com",
					Password: "this is password",
					Phone:    "12300001234",
					Ctime:    now,
				}).Return(nil)
				return d, c
			},
			ctx: context.Background(),
			id:  123,
			wantUser: domain.User{
				Id:       123,
				Email:    "123@qq.com",
				Password: "this is password",
				Phone:    "12300001234",
				Ctime:    now,
			},
			wantErr: nil,
		},
		{
			name: "缓存命中",
			mock: func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache) {
				// 缓存未命中，查了缓存，但是没结果
				c := cachemocks.NewMockUserCache(ctrl)
				c.EXPECT().Get(gomock.Any(), int64(123)).
					Return(domain.User{
						Id:       123,
						Email:    "123@qq.com",
						Password: "this is password",
						Phone:    "12300001234",
						Ctime:    now,
					}, nil)

				d := daomocks.NewMockUserDAO(ctrl)
				return d, c
			},
			ctx: context.Background(),
			id:  123,
			wantUser: domain.User{
				Id:       123,
				Email:    "123@qq.com",
				Password: "this is password",
				Phone:    "12300001234",
				Ctime:    now,
			},
			wantErr: nil,
		},
		{
			name: "缓存未命中，查询失败",
			mock: func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache) {
				// 缓存未命中，查了缓存，但是没结果
				c := cachemocks.NewMockUserCache(ctrl)
				c.EXPECT().Get(gomock.Any(), int64(123)).
					Return(domain.User{}, cache.ErrKeyNotExist)

				d := daomocks.NewMockUserDAO(ctrl)
				d.EXPECT().FindById(gomock.Any(), int64(123)).
					Return(dao.User{}, errors.New("mock db error"))
				return d, c
			},
			ctx:      context.Background(),
			id:       123,
			wantUser: domain.User{},
			wantErr:  errors.New("mock db error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ud, uc := tc.mock(ctrl)
			repo := NewUserRepository(ud, uc)
			u, err := repo.FindById(tc.ctx, tc.id)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
			time.Sleep(time.Second) // 用于走到 go routine 的路径
		})
	}
}

func TestCachedUserRepository_FindByPhone(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		c     context.Context
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			got, err := r.FindByPhone(tt.args.c, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCachedUserRepository_Update(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		c context.Context
		u domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			if err := r.Update(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCachedUserRepository_domainToEntity(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		u domain.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dao.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			if got := r.domainToEntity(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("domainToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCachedUserRepository_entityToDomain(t *testing.T) {
	type fields struct {
		dao   dao.UserDAO
		cache cache.UserCache
	}
	type args struct {
		u dao.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CachedUserRepository{
				dao:   tt.fields.dao,
				cache: tt.fields.cache,
			}
			if got := r.entityToDomain(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("entityToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		d dao.UserDAO
		c cache.UserCache
	}
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.d, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
