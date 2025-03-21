package bootstrap

import (
	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.Serve()
}
