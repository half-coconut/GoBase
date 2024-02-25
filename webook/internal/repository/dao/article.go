package dao

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ArticleDAO interface {
	Insert(ctx context.Context, art Article) (int64, error)
	UpdateById(ctx *gin.Context, article Article) error
}

func NewGORMArticleDAO(db *gorm.DB) ArticleDAO {
	return &GORMArticleDAO{
		db: db,
	}
}

type GORMArticleDAO struct {
	db *gorm.DB
}

func (dao *GORMArticleDAO) Insert(ctx context.Context, art Article) (int64, error) {
	now := time.Now().UnixMilli()
	art.Ctime = now
	art.Utime = now
	err := dao.db.WithContext(ctx).Create(&art).Error
	return art.Id, err
}

func (dao *GORMArticleDAO) UpdateById(ctx *gin.Context, art Article) error {
	now := time.Now().UnixMilli()
	art.Utime = now
	// 依赖 gorm 忽略零值的特性，默认用 id 主键进行更新，但是不建议，可读性很差，不清楚 updates 具体更新了啥
	//err := dao.db.WithContext(ctx).Updates(&art).Error
	res := dao.db.WithContext(ctx).Model(&art).Where("id=? AND author_id=?", art.Id, art.AuthorId).
		Updates(map[string]any{
			"title":   art.Title,
			"content": art.Content,
			"utime":   art.Utime,
		})
	if res.Error != nil {
		return res.Error
	}
	// 更新行数
	if res.RowsAffected == 0 {
		//return errors.New("更新失败，可能是作者对不上")
		return fmt.Errorf("更新失败，可能是作者对不上,id %d. author_id %d", art.Id, art.AuthorId)
	}
	return res.Error
}

// Article 这是制作库
type Article struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 长度 1024
	Title   string `gorm:"type=varchar(1024)"`
	Content string `gorm:"BLOG"`
	// 如何设计索引
	// 对于创作者来说，select * from article where author_id = 123.
	// 对于单独查询某一篇，select * from article where id = 1.
	// 实践：要按照创建/更新时间 倒序排序
	// select * from article where author_id = 123 order by `ctime` desc.
	//  - 在 author_id 和 ctime 上创建联合索引
	//  - 在 author_id 创建索引

	// 在 author_id 创建索引
	AuthorId int64 `gorm:"index"`
	//AuthorId int64 `gorm:"index=aid_ctime"`
	//Ctime    int64 `gorm:"index=aid_ctime"`
	Ctime int64
	Utime int64
}
