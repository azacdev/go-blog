package cmd

import (
	"github.com/azacdev/go-blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  `Database seeder`,
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}
