package view

import (
	"github.com/azacdev/go-blog/pkg/converters"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToURLValues(sessions.Flash(c, "old"))
	return data
}
