package repositories

import (
	ArticleModel "github.com/azacdev/go-blog/internal/modules/article/models"
	"github.com/azacdev/go-blog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Find(&articles)

	return articles
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}
