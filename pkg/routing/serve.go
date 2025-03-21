package routing

import (
	"fmt"
	"log"

	"github.com/azacdev/go-blog/internal/modules/home/routes"
	"github.com/azacdev/go-blog/pkg/config"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}

}

func RegisterRoutes() {
	routes.Routes(GetRouter())
}
