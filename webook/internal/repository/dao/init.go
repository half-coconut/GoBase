package dao

import "gorm.io/gorm"

// InitTable 初始化表结构
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
