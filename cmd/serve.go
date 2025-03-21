package cmd

import (
	"fmt"
	"net/http"

	"github.com/azacdev/go-blog/pkg/config"
	"github.com/azacdev/go-blog/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `Applications will be served on host and port defined in config yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func server() {
	config.Set()

	configs := config.Get()

	routes.Init()

	router := routes.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

}
