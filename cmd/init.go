package cmd

import (
	"fmt"
	"strings"

	"github.com/alimuddin7/skeleton-be/internal/generator"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  "Interactively generate a new production-ready Go microservice boilerplate.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fName, _ := cmd.Flags().GetString("name")
		fCode, _ := cmd.Flags().GetString("code")
		fType, _ := cmd.Flags().GetString("type")
		fDb, _ := cmd.Flags().GetString("db")

		projectName := fName
		if projectName == "" {
			projectName = "my-service"
		}
		serviceCode := fCode
		if serviceCode == "" {
			serviceCode = "01"
		}
		projectType := fType
		if projectType == "" {
			projectType = "Backend"
		}
		primaryDb := fDb
		if primaryDb == "" {
			primaryDb = "mysql"
		}

		var modules []string
		asynq := "No"
		grpc := "No"
		hostInput := "" // comma-separated host names

		var groups []*huh.Group

		// Step 1: Project Name
		if fName == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewInput().
					Title("Project Name").
					Description("Name of your Go microservice (e.g. payment-service)").
					Value(&projectName),
			))
		}

		// Step 2: Service Code
		if fCode == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewInput().
					Title("Service Code").
					Description("Project code and 2-digit identifier (e.g. OF01, OAG02)").
					Value(&serviceCode),
			))
		}

		// Step 3: Project Type
		if fType == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewSelect[string]().
					Title("Project Type").
					Options(huh.NewOptions("Backend", "Scheduler", "Worker", "Publisher", "gRPC")...).
					Value(&projectType),
			))
		}

		// Step 4: Primary Database
		if fDb == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewSelect[string]().
					Title("Primary Database").
					Options(huh.NewOptions("mysql", "postgresql")...).
					Value(&primaryDb),
			))
		}

		// Step 5: Additional Modules (optional — can skip by pressing Enter with nothing selected)
		groups = append(groups, huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Additional Modules").
				Description("Space to toggle, Enter to confirm. No selection = skip").
				Options(
					huh.NewOption("Redis Standalone", "redis"),
					huh.NewOption("Redis Cluster", "redis-cluster"),
					huh.NewOption("Kafka", "kafka"),
					huh.NewOption("NATS", "nats"),
					huh.NewOption("MinIO", "minio"),
				).
				Value(&modules),
		))

		// Step 6: External Hosts / API integration
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("External API Hosts").
				Description("Nama host eksternal, pisah koma jika lebih dari satu (e.g. core-payment,user-service). Kosongkan untuk skip").
				Value(&hostInput),
		))

		// Step 7: Asynq
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("Asynq (Background Queue)").
				Description("Redis-based async job queue").
				Options(huh.NewOptions("No", "Publisher", "Consumer", "Both")...).
				Value(&asynq),
		))

		// Step 8: gRPC
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("gRPC Support").
				Options(huh.NewOptions("No", "Server", "Client", "Both")...).
				Value(&grpc),
		))

		if err := huh.NewForm(groups...).Run(); err != nil {
			if err.Error() == "user aborted" {
				fmt.Println("Cancelled.")
				return nil
			}
			return fmt.Errorf("init cancelled: %w", err)
		}

		// Parse hosts
		var cleanHosts []string
		for _, h := range strings.Split(hostInput, ",") {
			h = strings.TrimSpace(h)
			if h != "" {
				cleanHosts = append(cleanHosts, strings.ToLower(h))
			}
		}

		config := generator.Config{
			ProjectName:  strings.TrimSpace(projectName),
			ServiceCode:  strings.TrimSpace(serviceCode),
			ProjectTypes: []string{projectType},
			Database:     primaryDb,
			Modules:      modules,
			Hosts:        cleanHosts,
		}

		if config.ProjectName == "" {
			return fmt.Errorf("project name cannot be empty")
		}

		// Add DB to modules
		if config.Database != "" {
			config.Modules = append(config.Modules, config.Database)
		}
		// Scheduler module
		for _, pt := range config.ProjectTypes {
			if pt == "Scheduler" {
				config.Modules = append(config.Modules, "scheduler")
			}
		}
		// Asynq modules
		if asynq != "No" {
			config.Modules = append(config.Modules, "redis-asynq")
			if asynq == "Publisher" || asynq == "Both" {
				config.Modules = append(config.Modules, "asynq-publisher")
			}
			if asynq == "Consumer" || asynq == "Both" {
				config.Modules = append(config.Modules, "asynq-consumer")
			}
		}
		// gRPC modules
		switch grpc {
		case "Server":
			config.Modules = append(config.Modules, "grpc-server")
		case "Client":
			config.Modules = append(config.Modules, "grpc-client")
		case "Both":
			config.Modules = append(config.Modules, "grpc-server", "grpc-client")
		}

		fmt.Printf("\nGenerating project \"%s\" [%s]...\n", config.ProjectName, config.ServiceCode)
		if err := generator.Generate(config.ProjectName, config); err != nil {
			return fmt.Errorf("generation failed: %w", err)
		}

		fmt.Println("✓ Project generated successfully!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().StringP("code", "c", "", "Service code (2 digits)")
	initCmd.Flags().StringP("db", "d", "", "Primary database (mysql, postgresql)")
	initCmd.Flags().StringP("type", "t", "", "Project type (Backend, Scheduler, Worker, Publisher, gRPC)")
}
