package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

/*
参数传结构体，不会触发内存逃逸
*/

func (dao *UserDAO) Insert(c context.Context, u User) error {
	// 存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	return dao.db.WithContext(c).Create(&u).Error // 这里是 gorm 的 db
}

// User 直接对应数据库表结构
// 别称：Entity、Model、PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一，设置为唯一索引
	Email    string `gorm:"unique"`
	Password string

	// 创建时间 毫秒数
	Ctime int64
	// 更新时间
	Utime int64
}
