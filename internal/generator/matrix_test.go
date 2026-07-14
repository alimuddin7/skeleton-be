package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratorMatrix(t *testing.T) {
	tests := []struct {
		name         string
		projectTypes []string
		database     string
		modules      []string
		hosts        []string
	}{
		{
			name:         "TC01_Backend_MySQL",
			projectTypes: []string{"Backend"},
			database:     "mysql",
			modules:      []string{"mysql"},
		},
		{
			name:         "TC02_Backend_PostgreSQL",
			projectTypes: []string{"Backend"},
			database:     "postgresql",
			modules:      []string{"postgresql"},
		},
		{
			name:         "TC03_Backend_NoDB",
			projectTypes: []string{"Backend"},
			database:     "",
			modules:      []string{},
		},
		{
			name:         "TC04_Scheduler_MySQL",
			projectTypes: []string{"Scheduler"},
			database:     "mysql",
			modules:      []string{"mysql", "scheduler"},
		},
		{
			name:         "TC05_Scheduler_NoDB",
			projectTypes: []string{"Scheduler"},
			database:     "",
			modules:      []string{"scheduler"},
		},
		{
			name:         "TC06_Worker_NATS_Consumer_NoDB",
			projectTypes: []string{"Worker"},
			database:     "",
			modules:      []string{"nats", "nats-consumer"},
		},
		{
			name:         "TC07_Publisher_Kafka_Publisher_MySQL",
			projectTypes: []string{"Publisher"},
			database:     "mysql",
			modules:      []string{"mysql", "kafka", "kafka-publisher"},
		},
		{
			name:         "TC08_gRPC_Server_MySQL",
			projectTypes: []string{"gRPC"},
			database:     "mysql",
			modules:      []string{"mysql", "grpc-server"},
		},
		{
			name:         "TC09_gRPC_Server_NoDB",
			projectTypes: []string{"gRPC"},
			database:     "",
			modules:      []string{"grpc-server"},
		},
		{
			name:         "TC10_Complex_Backend_Scheduler_Worker_Publisher_gRPC_Postgres_Redis_NatsBoth_Minio_Host",
			projectTypes: []string{"Backend", "Scheduler", "Worker", "Publisher", "gRPC"},
			database:     "postgresql",
			modules:      []string{"postgresql", "redis", "nats", "nats-consumer", "nats-publisher", "minio", "scheduler", "grpc-server", "grpc-client"},
			hosts:        []string{"payment-api"},
		},
		{
			name:         "TC11_Backend_RedisCluster",
			projectTypes: []string{"Backend"},
			database:     "",
			modules:      []string{"redis-cluster"},
		},
		{
			name:         "TC12_Worker_Asynq_Consumer_NoDB",
			projectTypes: []string{"Worker"},
			database:     "",
			modules:      []string{"asynq", "asynq-consumer"},
		},
	}

	fmt.Println("\n| Test Case | Project Types | Database | Modules | Hosts | Gen Status | Build Status | Error/Issues |")
	fmt.Println("|---|---|---|---|---|---|---|---|")

	for _, tc := range tests {
		tmpDir, err := os.MkdirTemp("", tc.name)
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		cfg := Config{
			ProjectName:  filepath.Base(tmpDir),
			ServiceCode:  "99",
			ProjectTypes: tc.projectTypes,
			Database:     tc.database,
			Modules:      tc.modules,
			Hosts:        tc.hosts,
		}

		genErr := Generate(tmpDir, cfg)
		genStatus := "Success"
		if genErr != nil {
			genStatus = "Failed"
		}

		buildStatus := "N/A"
		buildErrStr := ""
		if genErr == nil {
			// Try to build the generated code
			cmd := exec.Command("go", "build", "-o", filepath.Join(tmpDir, "app_bin"), filepath.Join(tmpDir, "main.go"))
			cmd.Dir = tmpDir
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				buildStatus = "Failed"
				buildErrLines := strings.Split(stderr.String(), "\n")
				if len(buildErrLines) > 0 {
					buildErrStr = buildErrLines[0] // just show the first line of error
					if len(buildErrLines) > 1 && buildErrLines[1] != "" {
						buildErrStr += "; " + buildErrLines[1]
					}
				}
			} else {
				buildStatus = "Success"
			}
		} else {
			buildErrStr = genErr.Error()
		}

		fmt.Printf("| %s | %v | %s | %v | %v | %s | %s | %s |\n",
			tc.name, tc.projectTypes, tc.database, tc.modules, tc.hosts, genStatus, buildStatus, buildErrStr)
	}
}
