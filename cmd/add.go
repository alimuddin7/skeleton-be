package cmd

import (
	"fmt"

	"github.com/alimuddin7/skeleton-be/internal/generator"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [module]",
	Short: "Add a new module to an existing project",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var module string
		if len(args) > 0 {
			module = args[0]
		} else {
			moduleSelect := promptui.Select{
				Label: "Select Module to Add",
				Items: []string{"mysql", "postgresql", "redis", "mongodb", "nats", "kafka", "host"},
			}
			_, module, _ = moduleSelect.Run()
		}

		fmt.Printf("Adding module %s...\n", module)
		err := generator.AddModule(module)
		if err != nil {
			fmt.Printf("Failed to add module: %v\n", err)
			return
		}

		fmt.Printf("Module %s added successfully!\n", module)
	},
}

var addRouteCmd = &cobra.Command{
	Use:   "route [name]",
	Short: "Add a new route/feature stack (Controller, Usecase, Repos, Models)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Generating feature stack for %s...\n", name)
		err := generator.AddFeature(name)
		if err != nil {
			fmt.Printf("Failed to generate feature: %v\n", err)
			return
		}
		fmt.Printf("Feature %s generated and registered successfully!\n", name)
	},
}

var addHostCmd = &cobra.Command{
	Use:   "host [name]",
	Short: "Add a new external host integration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Adding host integration for %s...\n", name)
		err := generator.AddHost(name)
		if err != nil {
			fmt.Printf("Failed to add host: %v\n", err)
			return
		}
		fmt.Printf("Host %s added successfully!\n", name)
	},
}

var addCrudCmd = &cobra.Command{
	Use:   "crud [name]",
	Short: "Add a new CRUD feature (Controller, Usecase, Repos, Models)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		dbType, _ := cmd.Flags().GetString("db")

		// If dbType is not provided, try to detect or ask
		if dbType == "" {
			dbSelect := promptui.Select{
				Label: "Select Database",
				Items: []string{"mysql", "postgresql"},
			}
			_, dbType, _ = dbSelect.Run()
		}

		fmt.Printf("Generating CRUD feature for %s using %s...\n", name, dbType)
		err := generator.AddCRUD(name, dbType)
		if err != nil {
			fmt.Printf("Failed to generate CRUD feature: %v\n", err)
			return
		}
		fmt.Printf("CRUD Feature %s generated and registered successfully!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addRouteCmd)
	addCmd.AddCommand(addHostCmd)
	addCmd.AddCommand(addCrudCmd)
	addCrudCmd.Flags().String("db", "", "Database type (mysql, postgresql)")

	addCmd.AddCommand(addHelperCmd)
}

var addHelperCmd = &cobra.Command{
	Use:   "helper [name]",
	Short: "Add a new helper function stub",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Generating helper %s...\n", name)
		err := generator.AddHelper(name)
		if err != nil {
			fmt.Printf("Failed to generate helper: %v\n", err)
			return
		}
		fmt.Printf("Helper %s generated successfully!\n", name)
	},
}
