package routes

import (
	"fmt"

	homeRoutes "github.com/azacdev/go-blog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	fmt.Println("Registering routes...")
	homeRoutes.Routes(router)
}
