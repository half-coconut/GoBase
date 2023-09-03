package ioc

import (
	"GoBase/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(webook-live-mysql:3308)/webook"), &gorm.Config{})
	if err != nil {
		// 在初始化过程中，panic
		// panic 使得 goroutine 直接结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
