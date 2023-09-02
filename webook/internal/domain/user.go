package domain

import "time"

// User 领域对象：是 DDD 中的聚合根 (Entity)
// 或者是 BO(Business Object)
type User struct {
	Id              int64
	Email           string
	Phone           string
	Password        string
	NickName        string    // 50个字符
	PersonalProfile string    // 200个字符
	Birthday        time.Time // 前端输入 1990-01-01 需要转化吗？
	Ctime           time.Time
}

// 不使用 *User 传值，会引发复制，使用指针是指向同一个内存地址
