package migration

import (
	"fmt"
	"log"

	articleModels "github.com/azacdev/go-blog/internal/modules/article/models"
	userModels "github.com/azacdev/go-blog/internal/modules/user/models"
	"github.com/azacdev/go-blog/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Failed to migrate")
	}

	fmt.Println("Migration done.")
}
