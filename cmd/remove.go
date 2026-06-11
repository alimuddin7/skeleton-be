package cmd

import (
	"fmt"

	"andromeda.ottopay.id/pt-rtsm-ottopay/skeleton-svc/internal/generator"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a module or feature from the project",
	Long:  "Remove domain components such as CRUD stacks, modules, routes, or helpers.",
}

var removeCrudCmd = &cobra.Command{
	Use:   "crud <name>",
	Short: "Remove a full CRUD stack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Removing CRUD stack for %s...\n", name)
		if err := generator.RemoveCRUD(name); err != nil {
			return fmt.Errorf("failed to remove CRUD: %w", err)
		}
		fmt.Printf("✓ CRUD feature %s removed successfully!\n", name)
		return nil
	},
}

var removeModuleCmd = &cobra.Command{
	Use:   "module <name>",
	Short: "Remove an infrastructure module",
	Long: `Remove an infrastructure module and its sub-modules from the project.
For messaging brokers, this also removes related consumer/publisher sub-modules.

Example:
  skeleton-be remove module redis
  skeleton-be remove module nats`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Removing module %s...\n", name)
		if err := generator.RemoveModule(name); err != nil {
			return fmt.Errorf("failed to remove module: %w", err)
		}
		fmt.Printf("✓ Module %s removed successfully!\n", name)
		return nil
	},
}

var removeRouteCmd = &cobra.Command{
	Use:   "route <name>",
	Short: "Remove a feature route stack",
	Long: `Remove a non-CRUD feature route and its generated files (controller, usecase, dto).

Example:
  skeleton-be remove route healthcheck`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Removing route %s...\n", name)
		if err := generator.RemoveRoute(name); err != nil {
			return fmt.Errorf("failed to remove route: %w", err)
		}
		fmt.Printf("✓ Route %s removed successfully!\n", name)
		return nil
	},
}

var removeHelperCmd = &cobra.Command{
	Use:   "helper <name>",
	Short: "Remove a helper utility",
	Long: `Remove a helper function file from the helpers/ directory.

Example:
  skeleton-be remove helper password-hash`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Removing helper %s...\n", name)
		if err := generator.RemoveHelper(name); err != nil {
			return fmt.Errorf("failed to remove helper: %w", err)
		}
		fmt.Printf("✓ Helper %s removed successfully!\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.AddCommand(removeModuleCmd)
	removeCmd.AddCommand(removeRouteCmd)
	removeCmd.AddCommand(removeCrudCmd)
	removeCmd.AddCommand(removeHelperCmd)
}

