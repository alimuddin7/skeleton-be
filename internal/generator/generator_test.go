package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateAndPreserve(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "skeleton-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	cfg := Config{
		ProjectName:  "test-service",
		ServiceCode:  "TS01",
		ProjectTypes: []string{"Backend"},
		Database:     "mysql",
		Modules:      []string{"mysql"},
	}

	// 1. First generation
	err = Generate(tmpDir, cfg)
	if err != nil {
		t.Fatalf("first generation failed: %v", err)
	}

	mainPath := filepath.Join(tmpDir, "main.go")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("main.go not generated: %v", err)
	}

	// Modify main.go
	modifiedContent := "// modified by user\npackage main\n"
	err = os.WriteFile(mainPath, []byte(modifiedContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Second generation (e.g. adding redis module)
	cfg.Modules = append(cfg.Modules, "redis")
	err = Generate(tmpDir, cfg)
	if err != nil {
		t.Fatalf("second generation failed: %v", err)
	}

	// Check if main.go was preserved
	content, err := os.ReadFile(mainPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != modifiedContent {
		t.Errorf("expected main.go content to be preserved, got:\n%s", string(content))
	}

	// Check that a main.go.new was created
	newPath := mainPath + ".new"
	if _, err := os.Stat(newPath); err != nil {
		t.Errorf("expected main.go.new to be created, but it doesn't exist: %v", err)
	}
}
