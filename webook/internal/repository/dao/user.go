package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
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

func (dao *UserDAO) FindById(c context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(c).First(&u, "id=?", id).Error
	return u, err
}

func (dao *UserDAO) FindByEmail(c context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email=?", email).First(&u).Error
	// 写法二
	//err := dao.db.WithContext(c).First(&u, "email=?", email).Error
	return u, err
}

func (dao *UserDAO) Insert(c context.Context, u User) error {
	// 存毫秒数
	now := time.Now().UnixMilli()
	// SELECT * FROM users where email=123@qq.com FOR UPDATE 间隙锁
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(c).Create(&u).Error // 这里是 gorm 的 db
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}
func (dao *UserDAO) Update(c context.Context, Id int64, nick_name, birthday, personal_profile string) (User, error) {
	var u User
	now := time.Now().UnixMilli()
	err := dao.db.Model(&u).WithContext(c).Where("id=?", Id).
		Update("NickName", nick_name).
		Update("Birthday", birthday).
		Update("PersonalProfile", personal_profile).Update("Utime", now).Error
	return u, err
}

// User 直接对应数据库表结构
// 别称：Entity、Model、PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一，设置为唯一索引
	Email           string `gorm:"unique"`
	Password        string
	NickName        string
	Birthday        string
	PersonalProfile string

	// 创建时间 毫秒数
	Ctime int64
	// 更新时间
	Utime int64
}
