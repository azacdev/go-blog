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

func (articleRepository *ArticleRepository) List(limit int) ([]ArticleModel.Article, error) {
	var articles []ArticleModel.Article

	result := articleRepository.DB.Limit(limit).Joins("User").Order("created_at DESC").Find(&articles)
	if result.Error != nil {
		return nil, result.Error // Return nil slice and the error
	}

	return articles, nil
}

func (articleRepository *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}

func (articleRepository *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}
