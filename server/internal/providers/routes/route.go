package routes

import (
	"fmt"

	articleRoutes "github.com/azacdev/go-blog/internal/modules/article/routes"
	homeRoutes "github.com/azacdev/go-blog/internal/modules/home/routes"
	userRoutes "github.com/azacdev/go-blog/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	fmt.Println("Registering routes...")
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
