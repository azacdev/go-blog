package view

import (
	"github.com/azacdev/go-blog/internal/modules/user/helpers"
	UserResponse "github.com/azacdev/go-blog/internal/modules/user/responses"
	"github.com/azacdev/go-blog/pkg/converters"
	"github.com/azacdev/go-blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	if data == nil {
		data = gin.H{}
	}

	user, err := helpers.Auth(c)

	if err != nil {
		// Handle error - either set empty user or log it
		data["AUTH"] = UserResponse.User{}
	} else {
		data["AUTH"] = user
	}

	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToURLValues(sessions.Flash(c, "old"))
	return data
}
