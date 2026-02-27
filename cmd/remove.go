package cmd

import (
	"fmt"

	"github.com/alimuddin7/skeleton-be/internal/generator"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a module or feature from the project",
	Long:  "Remove domain components such as CRUD stacks or other features.",
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

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.AddCommand(removeCrudCmd)
}
