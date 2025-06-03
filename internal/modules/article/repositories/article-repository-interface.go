package repositories

import ArticleModel "github.com/azacdev/go-blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) ([]ArticleModel.Article, error)
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
}
