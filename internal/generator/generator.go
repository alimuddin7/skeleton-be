package generator

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates
var templateFS embed.FS

type Config struct {
	ProjectName  string   `json:"projectName"`
	ServiceCode  string   `json:"serviceCode"`
	Database     string   `json:"database"`
	ProjectTypes []string `json:"projectTypes"`
	ProjectType  string   `json:"projectType,omitempty"` // Deprecated: use ProjectTypes
	Modules      []string `json:"modules"`
	Hosts        []string `json:"hosts"`
}

func Generate(destDir string, cfg Config) error {
	fmt.Printf("[DEBUG] Generate called with destDir=%s\n", destDir)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("[DEBUG] Failed to create project directory: %v\n", err)
		return fmt.Errorf("error creating project directory %s: %v", destDir, err)
	}
	fmt.Printf("[DEBUG] Project directory created: %s\n", destDir)

	if err := createDirectoryStructure(destDir, cfg); err != nil {
		return err
	}

	if err := saveProjectState(destDir, cfg); err != nil {
		return err
	}

	templates := getBaseTemplates()
	appendModuleTemplates(cfg, templates)

	for tmplPath, destPath := range templates {
		if err := renderAppTemplate(tmplPath, filepath.Join(destDir, destPath), cfg); err != nil {
			return err
		}
	}

	return nil
}

func AddModule(module string) error {
	cfg, err := loadProjectState("skeleton.json")
	if err != nil {
		return err
	}

	if module == "host" {
		// If user just types 'add host' without name, we should probably prompt or error
		// For now, let's assume they should use 'add host [name]'
		return fmt.Errorf("please use 'add host [name]' to add an external host")
	}

	if isModulePresent(cfg, module) {
		return nil
	}

	cfg.Modules = append(cfg.Modules, module)
	return Generate(".", *cfg)
}

func AddHost(name string) error {
	cfg, err := loadProjectState("skeleton.json")
	if err != nil {
		return err
	}

	lowerName := strings.ToLower(name)
	for _, h := range cfg.Hosts {
		if h == lowerName {
			return nil
		}
	}

	cfg.Hosts = append(cfg.Hosts, lowerName)

	// Create host directory
	hostDir := filepath.Join("hosts", lowerName)
	if err := os.MkdirAll(hostDir, 0755); err != nil {
		return err
	}

	// Render host template
	hostCfg := struct {
		ProjectName   string
		HostName      string
		HostNameLower string
		HostNameTitle string
	}{
		ProjectName:   cfg.ProjectName,
		HostName:      name,
		HostNameLower: lowerName,
		HostNameTitle: strings.Title(name),
	}

	if err := renderAppTemplate("templates/modules/host.go.tmpl", filepath.Join(hostDir, "host.go"), hostCfg); err != nil {
		return err
	}

	// Regenerate base files to include the new host (config, env, hosts/hosts.go)
	return Generate(".", *cfg)
}

func AddFeature(name string) error {
	cfg, err := loadProjectState("skeleton.json")
	if err != nil {
		return err
	}

	featureCfg := createFeatureConfig(cfg, name, false)
	if err := renderDomainComponents(featureCfg); err != nil {
		return err
	}

	return registerFeatureInBaseFiles(featureCfg.FeatureName, featureCfg.FeatureNameLower)
}

func AddCRUD(name, dbType string) error {
	cfg, err := loadProjectState("skeleton.json")
	if err != nil {
		return err
	}

	featureCfg := createFeatureConfig(cfg, name, true)

	// Create directories
	dirs := []string{
		"repositories",
		"models",
		fmt.Sprintf("controllers/v1/%s", featureCfg.FeatureNameLower),
		fmt.Sprintf("usecases/v1/%s", featureCfg.FeatureNameLower),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := renderDomainComponents(featureCfg); err != nil {
		return err
	}

	return registerCRUDInBaseFiles(featureCfg.FeatureName, featureCfg.FeatureNameLower)
}

func AddHelper(name string) error {
	_, err := loadProjectState("skeleton.json")
	if err != nil {
		return err
	}

	helperName := strings.ToLower(name)
	fileName := fmt.Sprintf("helpers/%s.go", helperName)

	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("helper %s already exists", fileName)
	}

	content := fmt.Sprintf(`package helpers

import "fmt"

// %s ...
func %s() {
	fmt.Println("Hello from %s helper")
}
`, strings.Title(helperName), strings.Title(helperName), strings.Title(helperName))

	return os.WriteFile(fileName, []byte(content), 0644)
}

func renderDomainComponents(cfg featureConfig) error {
	templates := map[string]string{
		"templates/domain/controller.go.tmpl": fmt.Sprintf("controllers/v1/%s/controller.go", cfg.FeatureNameLower),
		"templates/domain/usecase.go.tmpl":    fmt.Sprintf("usecases/v1/%s/usecase.go", cfg.FeatureNameLower),
		"templates/domain/repository.go.tmpl": fmt.Sprintf("repositories/%s.go", cfg.FeatureNameLower),
		"templates/domain/model.go.tmpl":      fmt.Sprintf("models/%s.go", cfg.FeatureNameLower),
	}

	for tmplPath, destPath := range templates {
		// Skip repository and model for non-CRUD features if desired
		if !cfg.IsCRUD && (strings.Contains(tmplPath, "repository") || strings.Contains(tmplPath, "model")) {
			continue
		}

		if err := renderAppTemplate(tmplPath, destPath, cfg); err != nil {
			return err
		}
	}
	return nil
}

func registerCRUDInBaseFiles(name, lower string) error {
	// Register in router
	if err := injectBelowMarker("routers/main.go", "// [V1_ROUTES_MARKER]",
		fmt.Sprintf("\t%sGroup := v1.Group(\"/%s\")\n\t%sGroup.Post(\"/\", ctrl.V1().%s().Create)\n\t%sGroup.Get(\"/:id\", ctrl.V1().%s().Get)\n\t%sGroup.Get(\"/\", ctrl.V1().%s().List)\n\t%sGroup.Put(\"/:id\", ctrl.V1().%s().Update)\n\t%sGroup.Delete(\"/:id\", ctrl.V1().%s().Delete)",
			lower, lower, lower, name, lower, name, lower, name, lower, name, lower, name)); err != nil {
		return err
	}

	// Register in V1 Controller Interface
	if err := injectBelowMarker("controllers/v1/controller.go", "// [V1_CONTROLLER_INTERFACE_MARKER]",
		fmt.Sprintf("\t%s() %sController", name, name)); err != nil {
		return err
	}

	// Register in V1 Usecase Interface
	if err := injectBelowMarker("usecases/v1/usecase.go", "// [V1_USECASE_INTERFACE_MARKER]",
		fmt.Sprintf("\t%s() %sUsecase", name, name)); err != nil {
		return err
	}

	return nil
}

// Internal helpers

func createDirectoryStructure(destDir string, cfg Config) error {
	// Base directories (always created)
	baseDirs := []string{
		"cmd", "configs", "constants", "controllers/v1", "usecases/v1",
		"routers", "models", "errorcodes", "docker", "migrations", "helpers", ".gitlab", ".gitlab/ci", ".gitlab/script",
	}

	for _, dir := range baseDirs {
		target := filepath.Join(destDir, dir)
		fmt.Printf("[DEBUG] Creating base directory: %s\n", target)
		if err := os.MkdirAll(target, 0755); err != nil {
			fmt.Printf("[DEBUG] Failed to create base directory %s: %v\n", target, err)
			return err
		}
	}

	// Module-specific directories
	moduleDirectories := map[string]string{
		"mysql":         "repositories/mysql",
		"postgresql":    "repositories/postgre",
		"redis":         "repositories/redis",
		"kafka":         "repositories/kafka",
		"nats":          "repositories/nats",
		"minio":         "repositories/minio",
		"redis-cluster": "repositories/redis_cluster",
		"grpc-server":   "grpc/server",
		"grpc-client":   "grpc/client",
		"scheduler":     "scheduler",
	}

	for _, module := range cfg.Modules {
		if dir, exists := moduleDirectories[module]; exists {
			if err := os.MkdirAll(filepath.Join(destDir, dir), 0755); err != nil {
				return err
			}
		}
	}

	// Create grpc/pb if any grpc module is present
	for _, mod := range cfg.Modules {
		if strings.HasPrefix(mod, "grpc-") {
			if err := os.MkdirAll(filepath.Join(destDir, "grpc/proto"), 0755); err != nil {
				return err
			}
			break
		}
	}

	// Host directories
	if len(cfg.Hosts) > 0 {
		if err := os.MkdirAll(filepath.Join(destDir, "hosts"), 0755); err != nil {
			return err
		}
		for _, host := range cfg.Hosts {
			if err := os.MkdirAll(filepath.Join(destDir, "hosts", host), 0755); err != nil {
				return err
			}
		}
	}

	return nil
}

func saveProjectState(destDir string, cfg Config) error {
	statePath := filepath.Join(destDir, "skeleton.json")
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(statePath, data, 0644)
}

func loadProjectState(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("missing skeleton.json: %v", err)
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	// Migrate ProjectType to ProjectTypes if needed
	if cfg.ProjectType != "" && len(cfg.ProjectTypes) == 0 {
		cfg.ProjectTypes = []string{cfg.ProjectType}
		cfg.ProjectType = "" // Clear it so it's not saved back in new format
	}

	return &cfg, nil
}

func isModulePresent(cfg *Config, module string) bool {
	for _, m := range cfg.Modules {
		if m == module {
			return true
		}
	}
	return false
}

func isProjectTypePresent(cfg Config, pType string) bool {
	for _, t := range cfg.ProjectTypes {
		if t == pType {
			return true
		}
	}
	return false
}

func getBaseTemplates() map[string]string {
	return map[string]string{
		"templates/base/main.go.tmpl":                               "main.go",
		"templates/base/go.mod.tmpl":                                "go.mod",
		"templates/base/constants/constants.go.tmpl":                "constants/general.go",
		"templates/base/helpers/models/models.go.tmpl":              "helpers/models/meta.go",
		"templates/base/helpers/models/log_models.go.tmpl":          "helpers/models/log.go",
		"templates/base/routers/router.go.tmpl":                     "routers/main.go",
		"templates/base/helpers/middleware.go.tmpl":                 "helpers/middleware.go",
		"templates/base/helpers/logger.go.tmpl":                     "helpers/logger.go",
		"templates/base/controllers/controller.go.tmpl":             "controllers/controller.go",
		"templates/base/controllers/v1/v1_controller.go.tmpl":       "controllers/v1/controller.go",
		"templates/base/usecases/v1/usecase.go.tmpl":                "usecases/v1/usecase.go",
		"templates/base/configs/config.go.tmpl":                     "configs/config.go",
		"templates/base/helpers/hc_helpers.go.tmpl":                 "helpers/health.go",
		"templates/base/helpers/utils.go.tmpl":                      "helpers/utils.go",
		"templates/base/helpers/auth.go.tmpl":                       "helpers/auth.go",
		"templates/base/helpers/http.go.tmpl":                       "helpers/http.go",
		"templates/base/infra/Makefile.tmpl":                        "Makefile",
		"templates/base/docker/Dockerfile.tmpl":                     "docker/Dockerfile",
		"templates/base/docker/docker-compose.yml.tmpl":             "docker/docker-compose.yml",
		"templates/base/docker/docker-compose-prod.yml.tmpl":        "docker/docker-compose-prod.yml",
		"templates/base/docker/docker-compose-stg.yml.tmpl":         "docker/docker-compose-stg.yml",
		"templates/base/configs/env.tmpl":                           ".env.example",
		"templates/base/infra/gitlab-ci.yml.tmpl":                   ".gitlab-ci.yml",
		"templates/base/gitlab/ci/build.gitlab-ci.yml.tmpl":         ".gitlab/ci/build.gitlab-ci.yml",
		"templates/base/gitlab/ci/deploy.gitlab-ci.yml.tmpl":        ".gitlab/ci/deploy.gitlab-ci.yml",
		"templates/base/gitlab/ci/sonar-scanner.gitlab-ci.yml.tmpl": ".gitlab/ci/sonar-scanner.gitlab-ci.yml",
		"templates/base/gitlab/script/command.sh.tmpl":              ".gitlab/script/command.sh",
		"templates/base/errorcodes/errorcodes.json.tmpl":            "errorcodes.json",
		"templates/base/errorcodes/errorcodes_en.json.tmpl":         "errorcodes/errorcodes-en.json",
	}
}

func appendModuleTemplates(cfg Config, templates map[string]string) {
	// Add databases.go.tmpl only if there are database modules
	hasDatabase := false
	for _, mod := range cfg.Modules {
		if mod == "mysql" || mod == "postgresql" || mod == "redis" || mod == "kafka" || mod == "nats" || mod == "minio" || mod == "redis-cluster" {
			hasDatabase = true
			break
		}
	}
	if hasDatabase {
		templates["templates/base/repositories/repositories.go.tmpl"] = "repositories/repository.go"
	}

	// Add scheduler.go.tmpl only if project type list contains Scheduler
	if isProjectTypePresent(cfg, "Scheduler") {
		templates["templates/base/scheduler/scheduler.go.tmpl"] = "scheduler/scheduler.go"
	}

	// Add hosts.go.tmpl only if there are hosts
	if len(cfg.Hosts) > 0 {
		templates["templates/base/hosts/hosts.go.tmpl"] = "hosts/hosts.go"
	}

	// Add module-specific templates
	mapping := map[string]string{
		"mysql":         "repositories/mysql/mysql.go",
		"postgresql":    "repositories/postgre/postgre.go",
		"redis":         "repositories/redis/redis.go",
		"kafka":         "repositories/kafka/kafka.go",
		"nats":          "repositories/nats/nats.go",
		"minio":         "repositories/minio/minio.go",
		"redis-cluster": "repositories/redis_cluster/redis_cluster.go",
		"grpc-server":   "grpc/server/server.go",
		"grpc-client":   "grpc/client/client.go",
	}
	for _, mod := range cfg.Modules {
		if dest, ok := mapping[mod]; ok {
			tmplName := strings.ReplaceAll(mod, "-", "_")
			if mod == "postgresql" {
				tmplName = "postgre"
			}
			templates[fmt.Sprintf("templates/modules/%s.go.tmpl", tmplName)] = dest
		}
	}

	// Add proto template if any grpc module is present
	for _, mod := range cfg.Modules {
		if strings.HasPrefix(mod, "grpc-") {
			templates["templates/modules/service.proto.tmpl"] = "grpc/proto/service.proto"
			break
		}
	}

	// Add host-specific templates
	for _, host := range cfg.Hosts {
		templates["templates/modules/host.go.tmpl"] = fmt.Sprintf("hosts/%s/%s.go", host, host)
	}
}

type featureConfig struct {
	ProjectName      string
	FeatureName      string
	FeatureNameLower string
	IsCRUD           bool
}

func createFeatureConfig(cfg *Config, name string, isCRUD bool) featureConfig {
	return featureConfig{
		ProjectName:      cfg.ProjectName,
		FeatureName:      strings.Title(name),
		FeatureNameLower: strings.ToLower(name),
		IsCRUD:           isCRUD,
	}
}

// Remove renderFeatureComponents as it is superseded by renderDomainComponents

func registerFeatureInBaseFiles(name, lower string) error {
	// Updated to point to the new location if we changed directory structure for basic features
	// If we use v1/name/controller.go, the import path in router might change or usage might change.
	// But Fiber router just calls the handler.
	// We need to ensure the Controller initialization in main/wire/controller.go is correct.

	// Assuming standard Clean Architecture with V1Controller struct in base/controllers/v1/controller.go
	// The original `feature` generator added methods to `v1Controller`.
	// The `crud` generator added interfaces.
	// We should probably standardize.
	// For "AddFeature" (Basic), we used to add method to v1Controller.
	// For "AddCRUD", we injected interface fields.

	// Let's keep `AddFeature` simple: Add method to `v1Controller`.
	// renderDomainComponents uses templates/domain/controller.go.tmpl

	if err := injectBelowMarker("routers/main.go", "// [V1_ROUTES_MARKER]", fmt.Sprintf("\tv1.Get(\"/%s\", ctrl.V1().%s)", lower, name)); err != nil {
		return err
	}
	// Note: Basic feature adds method to V1Controller interface/struct directly in templates?
	// No, the template defines `func (ctrl *v1Controller) {{.FeatureName}}`.
	// So we need to add the signature to the interface.

	if err := injectBelowMarker("controllers/v1/controller.go", "// [V1_CONTROLLER_INTERFACE_MARKER]", fmt.Sprintf("\t%s(c fiber.Ctx) error", name)); err != nil {
		return err
	}
	if err := injectBelowMarker("usecases/v1/usecase.go", "// [V1_USECASE_INTERFACE_MARKER]", fmt.Sprintf("\t%s(ctx context.Context) hModels.Response", name)); err != nil {
		return err
	}
	return nil
}

func injectBelowMarker(filePath, marker, code string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err // Could be missing if called prematurely
	}
	if bytes.Contains(content, []byte(code)) {
		return nil
	}
	newContent := bytes.Replace(content, []byte(marker), []byte(marker+"\n"+code), 1)
	return os.WriteFile(filePath, newContent, 0644)
}

func renderAppTemplate(tmplPath, destPath string, data interface{}) error {
	content, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		// Try to list the directory to see what's there
		dir := filepath.Dir(tmplPath)
		entries, _ := templateFS.ReadDir(dir)
		var files []string
		for _, e := range entries {
			files = append(files, e.Name())
		}
		fmt.Printf("Error reading embedded template %s: %v. Available files in %s: %v\n", tmplPath, err, dir, files)
		return fmt.Errorf("error reading embedded template %s: %v. Available files in %s: %v", tmplPath, err, dir, files)
	}

	tmpl, err := template.New(filepath.Base(tmplPath)).Funcs(template.FuncMap{
		"title": strings.Title,
		"upper": strings.ToUpper,
		"untitle": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToLower(s[0:1]) + s[1:]
		},
		"has": func(slice []string, item string) bool {
			for _, s := range slice {
				if s == item {
					return true
				}
			}
			return false
		},
	}).Parse(string(content))
	if err != nil {
		return fmt.Errorf("error parsing template %s: %v", tmplPath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("error executing template %s: %v", tmplPath, err)
	}

	dirToCreate := filepath.Dir(destPath)
	fmt.Printf("[DEBUG] renderAppTemplate: writing to %s, ensuring dir %s exists\n", destPath, dirToCreate)
	if err := os.MkdirAll(dirToCreate, 0755); err != nil {
		fmt.Printf("[DEBUG] Failed to create parent directory %s: %v\n", dirToCreate, err)
		return fmt.Errorf("error creating directory for %s: %v", destPath, err)
	}

	if err := os.WriteFile(destPath, buf.Bytes(), 0644); err != nil {
		fmt.Printf("[DEBUG] Failed to write file %s: %v\n", destPath, err)
		return fmt.Errorf("error writing to %s: %v", destPath, err)
	}
	fmt.Printf("[DEBUG] Successfully wrote file: %s\n", destPath)
	return nil
}
