package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skeleton-be",
	Short: "A dynamic boilerplate generator for Fiber v3 microservices",
	Long: `skeleton-be is a CLI tool that generates production-ready Go microservices
using Fiber v3, GORM v2, and Zerolog, based on clean architecture patterns.

Examples:
  skeleton-be init
  skeleton-be init --name my-service --code 01 --type Backend --db mysql
  skeleton-be add crud user
  skeleton-be add route notification
  skeleton-be add host core-payment`,
}

// Root exposes rootCmd so main.go can pass it to fang.Execute.
func Root() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().Bool("h", false, "display help")
}
