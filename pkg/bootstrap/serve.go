package bootstrap

import (
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/html"
	"github.com/azacdev/go-blog/pkg/routing"
	"github.com/azacdev/go-blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
