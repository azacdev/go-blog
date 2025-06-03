package bootstrap

import (
	seeder "github.com/azacdev/go-blog/internal/database/seeder"
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
