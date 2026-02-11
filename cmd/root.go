package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skeleton-be",
	Short: "A dynamic boilerplate generator for Fiber v3 microservices",
	Long: `skeleton-be is a CLI tool that generates production-ready Go microservices using Fiber v3, GORM v2, and Zerolog, based on clean architecture patterns.

Examples:
  # Initialize a new project with interactive prompts
  skeleton-be init

  # Initialize a new project with flags
  skeleton-be init --name my-service --code 01 --type Backend --db mysql

  # Add a new feature module (CRUD)
  skeleton-be add crud --name user --db mysql

  # Add a basic feature (Controller-Usecase only)
  skeleton-be add feature --name notification

  # Add a helper function
  skeleton-be add helper --name date_utils

  # Add a new external host integration
  skeleton-be add host --name core-payment

Available Commands:
  init        Initialize a new project
  add         Add a new module (crud, feature, helper, host)
  help        Help about any command`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Root flags if any
}
