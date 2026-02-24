package cmd

import (
	"fmt"

	"github.com/alimuddin7/skeleton-be/internal/generator"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a module or feature to an existing project",
	Long:  "Add infrastructure modules, feature routes, CRUD stacks, helpers, or host integrations.",
}

var addModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Add an infrastructure module (mysql, postgresql, redis, kafka, nats, minio)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		module := ""
		if len(args) > 0 {
			module = args[0]
		} else {
			if err := huh.NewSelect[string]().
				Title("Select Module").
				Options(huh.NewOptions("mysql", "postgresql", "redis", "redis-cluster", "kafka", "nats", "minio")...).
				Value(&module).
				Run(); err != nil {
				return fmt.Errorf("add module cancelled: %w", err)
			}
		}
		fmt.Printf("Adding module %s...\n", module)
		if err := generator.AddModule(module); err != nil {
			return fmt.Errorf("failed to add module: %w", err)
		}
		fmt.Printf("✓ Module %s added successfully!\n", module)
		return nil
	},
}

var addRouteCmd = &cobra.Command{
	Use:   "route <name>",
	Short: "Add a new route/feature stack (Controller, Usecase, Repository, Model)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Generating feature stack for %s...\n", name)
		if err := generator.AddFeature(name); err != nil {
			return fmt.Errorf("failed to generate feature: %w", err)
		}
		fmt.Printf("✓ Feature %s generated successfully!\n", name)
		return nil
	},
}

var addCrudCmd = &cobra.Command{
	Use:   "crud <name>",
	Short: "Add a CRUD feature stack (Controller, Usecase, Repository, Model)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		dbType, _ := cmd.Flags().GetString("db")
		if dbType == "" {
			if err := huh.NewSelect[string]().
				Title("Select Database").
				Options(huh.NewOptions("mysql", "postgresql")...).
				Value(&dbType).
				Run(); err != nil {
				return fmt.Errorf("add crud cancelled: %w", err)
			}
		}
		fmt.Printf("Generating CRUD for %s using %s...\n", name, dbType)
		if err := generator.AddCRUD(name, dbType); err != nil {
			return fmt.Errorf("failed to generate CRUD: %w", err)
		}
		fmt.Printf("✓ CRUD feature %s generated successfully!\n", name)
		return nil
	},
}

var addHostCmd = &cobra.Command{
	Use:   "host <name>",
	Short: "Add a new external host integration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Adding host integration for %s...\n", name)
		if err := generator.AddHost(name); err != nil {
			return fmt.Errorf("failed to add host: %w", err)
		}
		fmt.Printf("✓ Host %s added successfully!\n", name)
		return nil
	},
}

var addHelperCmd = &cobra.Command{
	Use:   "helper <name>",
	Short: "Add a new helper function stub",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Printf("Generating helper %s...\n", name)
		if err := generator.AddHelper(name); err != nil {
			return fmt.Errorf("failed to generate helper: %w", err)
		}
		fmt.Printf("✓ Helper %s generated successfully!\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addModuleCmd)
	addCmd.AddCommand(addRouteCmd)
	addCmd.AddCommand(addCrudCmd)
	addCmd.AddCommand(addHostCmd)
	addCmd.AddCommand(addHelperCmd)
	addCrudCmd.Flags().StringP("db", "d", "", "Database type (mysql, postgresql)")
}
