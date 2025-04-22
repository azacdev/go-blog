package services

import (
	"errors"

	ArticleModel "github.com/azacdev/go-blog/internal/modules/article/models"
	ArticleRepository "github.com/azacdev/go-blog/internal/modules/article/repositories"
	"github.com/azacdev/go-blog/internal/modules/article/request/articles"
	ArticleResponse "github.com/azacdev/go-blog/internal/modules/article/responses"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)

	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return response, errors.New("error in creating the article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}
