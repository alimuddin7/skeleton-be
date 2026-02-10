package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skeleton-be",
	Short: "A dynamic boilerplate generator for Fiber v3 microservices",
	Long:  `skeleton-be is a CLI tool that generates production-ready Go microservices using Fiber v3, GORM v2, and Zerolog, based on skeleton-svc patterns.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root flags if any
}
