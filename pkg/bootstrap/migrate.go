package bootstrap

import (
	migration "github.com/azacdev/go-blog/internal/database/migration"
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
