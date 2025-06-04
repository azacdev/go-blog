package bootstrap

import (
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/database"
	"github.com/azacdev/go-blog/pkg/routing"
	"github.com/azacdev/go-blog/pkg/sessions"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()
	sessions.Start(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
