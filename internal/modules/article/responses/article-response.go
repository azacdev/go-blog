package responses

import (
	"fmt"

	articleModels "github.com/azacdev/go-blog/internal/modules/article/models"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article articleModels.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     "https://plus.unsplash.com/premium_photo-1721268770804-f9db0ce102f8?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []articleModels.Article) Articles {
	var response Articles

	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}

	return response
}
