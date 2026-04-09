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
	Short: "Initialize a new microservice project",
	Long: `Initialize a new production-ready Go microservice boilerplate through an interactive 9-step wizard.

Available Project Types:
- Backend: REST API using Fiber v3
- Scheduler: Background cron jobs using robfig/cron
- Worker: Message consumer for NATS/Kafka/Asynq
- Publisher: Message producer for NATS/Kafka/Asynq
- gRPC: High-performance RPC (Server/Client)

Available Infrastructure Modules:
- Databases: MySQL, PostgreSQL
- Caching: Redis Standalone, Redis Cluster
- Messaging: NATS JetStream, Kafka, Asynq (Redis-based)
- Storage: MinIO`,
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
		projectTypes := make([]string, 0)
		var messagingBroker string
		messagingRole := "Consumer"

		projectType := fType
		if projectType == "" {
			projectType = "Backend"
		}
		primaryDb := fDb
		if primaryDb == "" {
			primaryDb = "mysql"
		}
		if primaryDb == "postgres" {
			primaryDb = "postgresql"
		}

		modules, _ := cmd.Flags().GetStringSlice("modules")
		grpc := "No"
		hostInput, _ := cmd.Flags().GetString("hosts")

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

		// Step 3: Project Type (multi-select)
		if fType == "" {
			groups = append(groups, huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Project Type").
					Description("Space to toggle, Enter to confirm. Select all that apply.").
					Options(
						huh.NewOption("Backend (REST API)", "Backend"),
						huh.NewOption("Scheduler (Cron Jobs)", "Scheduler"),
						huh.NewOption("Worker (Message Consumer)", "Worker"),
						huh.NewOption("Publisher (Message Producer)", "Publisher"),
						huh.NewOption("gRPC", "gRPC"),
					).
					Value(&projectTypes).
					Validate(func(t []string) error {
						if len(t) == 0 {
							return fmt.Errorf("please select at least one project type using Spacebar")
						}
						return nil
					}),
			))
		} else {
			projectTypes = []string{projectType}
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

		// Step 5: Additional Infrastructure Modules
		if !cmd.Flags().Changed("modules") {
			groups = append(groups, huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Additional Infrastructure Modules").
					Description("Space to toggle, Enter to confirm. No selection = skip").
					Options(
						huh.NewOption("Redis Standalone", "redis"),
						huh.NewOption("Redis Cluster", "redis-cluster"),
						huh.NewOption("MinIO", "minio"),
					).
					Value(&modules),
			))
		}

		// Step 6: Messaging Modules (NATS / Kafka / Asynq)
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("Messaging / Queue Broker").
				Description("Select a primary messaging broker to integrate.").
				Options(
					huh.NewOption("None", ""),
					huh.NewOption("NATS JetStream", "nats"),
					huh.NewOption("Kafka", "kafka"),
					huh.NewOption("Asynq (Redis-based)", "asynq"),
				).
				Value(&messagingBroker),
		))

		// Step 7: Role for messaging — always shown, ignored if no broker selected
		groups = append(groups, huh.NewGroup(
			huh.NewSelect[string]().
				Title("Messaging Role").
				Description("How will this service use the messaging modules? (Ignored if no broker selected above)").
				Options(huh.NewOptions("Consumer", "Publisher", "Both")...).
				Value(&messagingRole),
		))

		// Step 8: External Hosts / API integration
		if !cmd.Flags().Changed("hosts") {
			groups = append(groups, huh.NewGroup(
				huh.NewInput().
					Title("External API Hosts").
					Description("Nama host eksternal, pisah koma jika lebih dari satu (e.g. core-payment,user-service). Kosongkan untuk skip").
					Value(&hostInput),
			))
		}

		// Step 9: gRPC
		fGrpc, _ := cmd.Flags().GetString("grpc")
		if fGrpc != "" {
			grpc = fGrpc
		} else {
			groups = append(groups, huh.NewGroup(
				huh.NewSelect[string]().
					Title("gRPC Support").
					Options(huh.NewOptions("No", "Server", "Client", "Both")...).
					Value(&grpc),
			))
		}

		if len(groups) > 0 {
			if err := huh.NewForm(groups...).Run(); err != nil {
				if err.Error() == "user aborted" {
					fmt.Println("Cancelled.")
					return nil
				}
				return fmt.Errorf("init cancelled: %w", err)
			}
		}

		// Parse hosts
		var cleanHosts []string
		for _, h := range strings.Split(hostInput, ",") {
			h = strings.TrimSpace(h)
			if h != "" {
				cleanHosts = append(cleanHosts, strings.ToLower(h))
			}
		}

		// If fType was provided via flag (not interactive), use it as-is
		if fType != "" {
			projectTypes = []string{fType}
		}

		// [DEBUG] Log selections
		fmt.Printf("\n--- Debug Selections ---\n")
		fmt.Printf("ProjectTypes: %v\n", projectTypes)
		fmt.Printf("Modules (Infrastructure): %v\n", modules)
		fmt.Printf("MessagingBroker: %s\n", messagingBroker)
		fmt.Printf("MessagingRole: %s\n", messagingRole)
		fmt.Printf("------------------------\n")

		config := generator.Config{
			ProjectName:  strings.TrimSpace(projectName),
			ServiceCode:  strings.TrimSpace(serviceCode),
			ProjectTypes: projectTypes,
			Database:     primaryDb,
			Modules:      modules,
			Hosts:        cleanHosts,
		}

		if config.ProjectName == "" {
			return fmt.Errorf("project name cannot be empty")
		}

		// Add primary DB to modules
		if config.Database != "" {
			config.Modules = append(config.Modules, config.Database)
		}

		// Scheduler module
		for _, pt := range config.ProjectTypes {
			if pt == "Scheduler" {
				config.Modules = append(config.Modules, "scheduler")
			}
		}

		// Add messaging broker and its role
		if messagingBroker != "" {
			// Always add the base infra module
			config.Modules = append(config.Modules, messagingBroker)
			// Add role-specific module strings
			if messagingRole == "Consumer" || messagingRole == "Both" {
				config.Modules = append(config.Modules, messagingBroker+"-consumer")
			}
			if messagingRole == "Publisher" || messagingRole == "Both" {
				config.Modules = append(config.Modules, messagingBroker+"-publisher")
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

		fmt.Printf("\n✓ Project \"%s\" generated and synchronized successfully!\n", config.ProjectName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().StringP("code", "c", "", "Service code (2 digits)")
	initCmd.Flags().StringP("db", "d", "", "Primary database (mysql, postgresql)")
	initCmd.Flags().StringP("type", "t", "", "Project type (Backend, Scheduler, Worker, Publisher, gRPC)")
	initCmd.Flags().StringSliceP("modules", "m", []string{}, "Additional modules (redis, kafka, nats, minio)")
	initCmd.Flags().StringP("hosts", "H", "", "External API hosts (comma separated)")
	initCmd.Flags().StringP("grpc", "g", "", "gRPC mode (No, Server, Client, Both)")
}
