package domain

// User 领域对象：是 DDD 中的聚合根 (Entity)
// 或者是 BO(Business Object)
type User struct {
	Email    string
	Password string
}

// 不使用 *User 传值，会引发复制，使用指针是指向同一个内存地址
