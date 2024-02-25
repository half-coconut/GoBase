package repository

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository/dao"
	"context"
	"github.com/gin-gonic/gin"
)

type ArticleRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx *gin.Context, art domain.Article) error
	//FindById(ctx *gin.Context, Id int64) domain.Article
}

type CachedArticleRepository struct {
	dao dao.ArticleDAO
}

func NewArticleRepository(d dao.ArticleDAO) ArticleRepository {
	return &CachedArticleRepository{
		dao: d,
	}
}

func (c *CachedArticleRepository) Create(ctx context.Context, art domain.Article) (int64, error) {
	return c.dao.Insert(ctx, dao.Article{
		Title:    art.Title,
		Content:  art.Content,
		AuthorId: art.Author.Id,
	})
}

func (c *CachedArticleRepository) Update(ctx *gin.Context, art domain.Article) error {
	return c.dao.UpdateById(ctx, dao.Article{
		Id:       art.Id,
		Title:    art.Title,
		Content:  art.Content,
		AuthorId: art.Author.Id,
	})
}

//func (c *CachedArticleRepository) FindById(ctx *gin.Context, Id int64) domain.Article {
//	//TODO implement me
//	panic("implement me")
//}
