package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration tools",
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		timestamp := time.Now().Format("20060102150405")

		upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, name)
		downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, name)

		migrationsDir := "migrations"
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			os.Mkdir(migrationsDir, 0755)
		}

		createFile(filepath.Join(migrationsDir, upFileName))
		createFile(filepath.Join(migrationsDir, downFileName))

		fmt.Printf("Created migration files:\n%s\n%s\n", upFileName, downFileName)
	},
}

func createFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", path, err)
		return
	}
	defer f.Close()
	f.WriteString("-- Write your migration here\n")
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateCreateCmd)
	// TODO: Add up/down commands integrating with golang-migrate or similar
}
