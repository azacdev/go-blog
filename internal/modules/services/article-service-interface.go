package services

import (
	"github.com/azacdev/go-blog/internal/modules/article/request/articles"
	ArticleResponse "github.com/azacdev/go-blog/internal/modules/article/responses"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() (ArticleResponse.Articles, error)
	GetStoriesArticles() (ArticleResponse.Articles, error)
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
