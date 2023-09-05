package service

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository"
	repomocks "GoBase/webook/internal/repository/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"testing"
	"time"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		repo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Edit(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
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
			svc := &userService{
				repo: tt.fields.repo,
			}
			if err := svc.Edit(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Edit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_FindOrCreate(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
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
			svc := &userService{
				repo: tt.fields.repo,
			}
			got, err := svc.FindOrCreate(tt.args.c, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Login(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) repository.UserRepository
		email    string
		password string
		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Phone:    "13800001234",
						Password: "$2a$10$M2kflFNF.36eM8IGz2Ig3e1ddad.xfmzIe9Lf2m5fF.Frba7gtrGW",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "hello@123",
			wantUser: domain.User{
				Email:    "123@qq.com",
				Phone:    "13800001234",
				Password: "$2a$10$M2kflFNF.36eM8IGz2Ig3e1ddad.xfmzIe9Lf2m5fF.Frba7gtrGW",
				Ctime:    now,
			},
			wantErr: nil,
		},
		{
			name: "用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, repository.ErrUserNotFound)
				return repo
			},
			email:    "123@qq.com",
			password: "hello@123",

			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "DB 错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, errors.New("mock db 错误"))
				return repo
			},
			email:    "123@qq.com",
			password: "hello@123",

			wantUser: domain.User{},
			wantErr:  errors.New("mock db 错误"),
		},
		{
			name: "密码不对",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Phone:    "13800001234",
						Password: "$2a$10$M2kflFNF.36eM8IGz2Ig3e1ddad.xfmzIe9Lf2m5fF.Frba7gtrGW",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123@qq.com",
			password: "hello@123456",
			wantUser: domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := NewUserService(tc.mock(ctrl))
			u, err := svc.Login(context.Background(), tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}

func Test_userService_Profile(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
	}
	type args struct {
		c  context.Context
		id int64
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
			svc := &userService{
				repo: tt.fields.repo,
			}
			got, err := svc.Profile(tt.args.c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Profile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Profile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_SignUp(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
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
			svc := &userService{
				repo: tt.fields.repo,
			}
			if err := svc.SignUp(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_UpdateNonSensitiveInfo(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
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
			svc := &userService{
				repo: tt.fields.repo,
			}
			if err := svc.UpdateNonSensitiveInfo(tt.args.c, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UpdateNonSensitiveInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncrypted(t *testing.T) {
	res, err := bcrypt.GenerateFromPassword([]byte("hello@123"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(res))
	}
}
