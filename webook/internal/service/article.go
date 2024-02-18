package service

import (
	"GoBase/webook/internal/domain"
	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	Save(ctx *gin.Context, art domain.Article) (int64, error)
}

type articleService struct {
}

func NewArticleService() ArticleService {
	return &articleService{}
}

func (a *articleService) Save(ctx *gin.Context, art domain.Article) (int64, error) {
	return 1, nil
}
