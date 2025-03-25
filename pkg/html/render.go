package html

import (
	"github.com/azacdev/go-blog/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	view.WithGlobalData(data)
	c.HTML(code, name, data)
}
