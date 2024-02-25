package service

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/repository"
	"errors"
	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	Save(ctx *gin.Context, art domain.Article) (int64, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (a *articleService) Save(ctx *gin.Context, art domain.Article) (int64, error) {
	if art.Id > 0 {
		err := a.repo.Update(ctx, art)
		return art.Id, err
	}
	return a.repo.Create(ctx, art)
}

func (a *articleService) update(ctx *gin.Context, art domain.Article) error {
	// 只要 author 不允许更新
	// 但是性能比较差...
	// 更优写法：
	artInDB := a.repo.FindById(ctx, art.Id)
	if art.Author.Id != artInDB.Author.Id {
		return errors.New("更新别人的数据失败")
	}
	return a.repo.Update(ctx, art)
}
