package cmd

import (
	"fmt"

	"github.com/alimuddin7/skeleton-be/internal/generator"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")
		serviceCode, _ := cmd.Flags().GetString("code")
		if serviceCode == "" {
			serviceCode, _ = cmd.Flags().GetString("service-code")
		}
		database, _ := cmd.Flags().GetString("db")
		if database == "" {
			database, _ = cmd.Flags().GetString("database")
		}
		projectType, _ := cmd.Flags().GetString("type")
		if projectType == "" {
			projectType, _ = cmd.Flags().GetString("project-type")
		}

		if projectName == "" {
			projectNamePrompt := promptui.Prompt{
				Label:   "Project Name",
				Default: "my-service",
			}
			projectName, _ = projectNamePrompt.Run()
		}

		if serviceCode == "" {
			serviceCodePrompt := promptui.Prompt{
				Label:   "Service Code (2 digits)",
				Default: "01",
			}
			serviceCode, _ = serviceCodePrompt.Run()
		}

		if projectType == "" {
			typeSelect := promptui.Select{
				Label: "Select Project Type",
				Items: []string{"Backend", "Scheduler", "Worker", "Publisher", "gRPC"},
			}
			_, projectType, _ = typeSelect.Run()
		}

		config := generator.Config{
			ProjectName:  projectName,
			ServiceCode:  serviceCode,
			ProjectTypes: []string{projectType},
			Modules:      []string{},
			Hosts:        []string{},
		}

		// Add scheduler module if project type list contains Scheduler
		for _, pt := range config.ProjectTypes {
			if pt == "Scheduler" {
				config.Modules = append(config.Modules, "scheduler")
			}
		}

		// Interactive module selection
		fmt.Println("\n=== Module Selection ===")

		modules := []struct {
			name  string
			label string
		}{
			{"mysql", "Should MySQL repository functionality be included?"},
			{"postgresql", "Should PostgreSQL repository functionality be included?"},
			{"redis", "Should Redis standalone functionality be included?"},
			{"redis-cluster", "Should Redis cluster functionality be included?"},
			{"kafka", "Should Kafka functionality be included?"},
			{"nats", "Should NATS functionality be included?"},
			{"minio", "Should MinIO functionality be included?"},
		}

		for _, mod := range modules {
			prompt := promptui.Select{
				Label: mod.label,
				Items: []string{"Yes", "No"},
			}
			_, result, _ := prompt.Run()
			if result == "Yes" {
				config.Modules = append(config.Modules, mod.name)
			}
		}

		// Ask about Redis Asynq (for queue/worker functionality)
		asynqPrompt := promptui.Select{
			Label: "Should Redis Asynq (queue/worker) functionality be included?",
			Items: []string{"No", "Publisher", "Consumer", "Both"},
		}
		_, asynqResult, _ := asynqPrompt.Run()
		if asynqResult != "No" {
			config.Modules = append(config.Modules, "redis-asynq")
			// Store asynq mode in config for template usage
			if asynqResult == "Publisher" || asynqResult == "Both" {
				config.Modules = append(config.Modules, "asynq-publisher")
			}
			if asynqResult == "Consumer" || asynqResult == "Both" {
				config.Modules = append(config.Modules, "asynq-consumer")
			}
		}

		// Ask about gRPC
		grpcPrompt := promptui.Select{
			Label: "Should gRPC functionality be included?",
			Items: []string{"None", "Server", "Client", "Both"},
		}
		_, grpcResult, _ := grpcPrompt.Run()
		switch grpcResult {
		case "Server":
			config.Modules = append(config.Modules, "grpc-server")
		case "Client":
			config.Modules = append(config.Modules, "grpc-client")
		case "Both":
			config.Modules = append(config.Modules, "grpc-server", "grpc-client")
		}

		fmt.Printf("\nGenerating project %s...\n", projectName)
		if err := generator.Generate(config.ProjectName, config); err != nil {
			fmt.Printf("Generation failed: %v\n", err)
			return
		}

		fmt.Println("Project generated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().StringP("code", "c", "", "Service code (2 digits)")
	initCmd.Flags().String("service-code", "", "Service code (alias for --code)")
	initCmd.Flags().StringP("db", "d", "", "Primary database (mysql, postgresql)")
	initCmd.Flags().StringP("type", "t", "", "Project type (Backend, Scheduler, Worker, Publisher, gRPC)")
	initCmd.Flags().String("database", "", "Primary database (alias for --db)")
	initCmd.Flags().String("project-type", "", "Project type (alias for --type)")
}
