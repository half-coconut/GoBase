package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicate = errors.New("用户邮箱或者手机号冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

type UserDAO interface {
	FindById(c context.Context, id int64) (User, error)
	FindByEmail(c context.Context, email string) (User, error)
	FindByPhone(c context.Context, phone string) (User, error)
	FindByWechat(c context.Context, openID string) (User, error)
	Insert(c context.Context, u User) error
	Update(c context.Context, Id int64, nick_name, birthday, personal_profile string) (User, error)
	UpdateNonZeroFields(c context.Context, u User) error
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

/*
参数传结构体，不会触发内存逃逸
*/

func (dao *GORMUserDAO) FindById(c context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(c).First(&u, "id=?", id).Error
	return u, err
}
func (dao *GORMUserDAO) FindByWechat(c context.Context, openID string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("wechat_open_id=?", openID).First(&u).Error
	// 写法二
	//err := dao.db.WithContext(c).First(&u, "email=?", email).Error
	return u, err
}

func (dao *GORMUserDAO) FindByEmail(c context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email=?", email).First(&u).Error
	// 写法二
	//err := dao.db.WithContext(c).First(&u, "email=?", email).Error
	return u, err
}
func (dao *GORMUserDAO) FindByPhone(c context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("phone=?", phone).First(&u).Error
	return u, err
}

func (dao *GORMUserDAO) Insert(c context.Context, u User) error {
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
			return ErrUserDuplicate
		}
	}
	return err
}
func (dao *GORMUserDAO) Update(c context.Context, Id int64, nick_name, birthday, personal_profile string) (User, error) {
	var u User
	now := time.Now().UnixMilli()
	err := dao.db.Model(&u).WithContext(c).Where("id=?", Id).
		Update("NickName", nick_name).
		Update("Birthday", birthday).
		Update("PersonalProfile", personal_profile).Update("Utime", now).Error
	return u, err
}

func (dao *GORMUserDAO) UpdateNonZeroFields(c context.Context, u User) error {
	// 这种写法是很不清晰的，因为它依赖了 gorm 的两个默认语义
	// 会使用 ID 来作为 WHERE 条件
	// 会使用非零值来更新
	// 另外一种做法是显式指定只更新必要的字段，
	// 那么这意味着 DAO 和 service 中非敏感字段语义耦合了
	return dao.db.Updates(&u).Error
}

// User 直接对应数据库表结构
// 别称：Entity、Model、PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一，设置为唯一索引
	Email    sql.NullString `gorm:"unique"`
	Password string

	// 唯一索引允许有多个空值
	// 但是不能有多个 ""
	Phone    sql.NullString `gorm:"unique"`
	NickName sql.NullString
	Birthday sql.NullInt64
	// 自我介绍
	// 指定是 varchar 这个类型，并且长度是 1024
	// 因此你可以看到在 web 里面有这个校验
	PersonalProfile sql.NullString `gorm:"type=varchar(1024)"`

	// 索引的最左匹配原则
	// 假如索引在 <A，B，C> 建好了
	// where 里面带了 ABC, 可以用，例如：A、AB、ABC 都能用
	// where 里面没有 A, 就不能用

	// 如果要创建联合索引，<unionid, openid>，用 openid 查询的时候不会走索引
	// <openid, unionid> 用 unionid 查询时，不会走索引
	// 微信绑定的字段
	WechatUnionID sql.NullString
	WechatOpenID  sql.NullString `gorm:"unique"` // 授权用户唯一标识

	// 创建时间 毫秒数
	Ctime int64
	// 更新时间
	Utime int64
}
