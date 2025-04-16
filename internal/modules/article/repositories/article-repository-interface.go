package repositories

import ArticleModel "github.com/azacdev/go-blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
}
