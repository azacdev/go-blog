package routes

import (
	homeRoutes "github.com/azacdev/go-blog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	homeRoutes.Routes(router)
}
