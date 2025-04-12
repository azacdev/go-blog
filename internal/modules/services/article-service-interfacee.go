package services

import ArticleResponse "github.com/azacdev/go-blog/internal/modules/article/responses"

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
}
