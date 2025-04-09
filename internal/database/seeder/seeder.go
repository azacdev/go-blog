package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/azacdev/go-blog/internal/modules/article/models"
	userModel "github.com/azacdev/go-blog/internal/modules/user/models"
	"github.com/azacdev/go-blog/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)

	if err != nil {
		log.Fatal("Failed to hash password")
		return
	}
	user := userModel.User{Name: "Azacdev", Email: "azacdev@gmail.com", Password: string(hashedPassword)}

	db.Create(&user)

	log.Printf("User created successfully with email address %s \n", user.Email)

	for i := 1; i < 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random Title %d", i), Content: fmt.Sprintf("I am in help please help me %d", i), UserID: 1}

		db.Create(&article)

		log.Printf("Article created successfully %s \n", article.Title)

	}

	log.Println("Seeder none...")
}
