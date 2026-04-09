package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "skeleton-be",
	Short: "Interactive boilerplate generator for Fiber v3 microservices",
	Long: `Skeleton-BE is a powerful CLI tool designed to scaffold production-ready Go microservices.
It follows Clean Architecture principles and integrates best-in-class libraries like Fiber v3, GORM v2, and Zerolog.

Features:
- Standardized project structure
- Plug-and-play infrastructure modules (DBs, Caches, Brokers)
- Multi-messaging role support (Consumer/Publisher)
- Automated CRUD and feature generation
- Docker and GitLab CI/CD integration

Usage Examples:
  skeleton-be init                         # Launch interactive wizard
  skeleton-be add module redis             # Add Redis standalone
  skeleton-be add module nats              # Add NATS with role selection
  skeleton-be add crud product             # Generate full CRUD for product entity
  skeleton-be add route health             # Generate a simple healthcheck route
  skeleton-be add host user-api            # Scaffold an external API client`,
}

// Root exposes rootCmd so main.go can pass it to fang.Execute.
func Root() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().Bool("h", false, "display help")
}
