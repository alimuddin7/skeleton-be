package cmd

import (
	"fmt"
	"skeleton-be/internal/generator"

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

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(addRouteCmd)
	addCmd.AddCommand(addHostCmd)
}
