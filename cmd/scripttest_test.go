package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func TestScripts(t *testing.T) {
	// Skip on Windows - test scripts use sh -c which requires Unix shell
	if runtime.GOOS == "windows" {
		t.Skip("scripttest uses Unix shell commands (sh -c), skipping on Windows")
	}

	// Build the kai binary
	exeName := "kai"
	binDir := t.TempDir()
	exe := filepath.Join(binDir, exeName)

	// Build from the parent directory (where main.go is)
	mainDir := filepath.Join("..", ".")
	if err := exec.Command("go", "build", "-o", exe, mainDir).Run(); err != nil {
		t.Fatalf("failed to build kai binary: %v", err)
	}

	// Create minimal engine with default commands plus kai
	timeout := 2 * time.Second
	engine := script.NewEngine()
	engine.Cmds["kai"] = script.Program(exe, nil, timeout)

	// Add binDir to PATH so 'sh -c kai ...' works in test scripts
	currentPath := os.Getenv("PATH")
	env := []string{"PATH=" + binDir + ":" + currentPath}

	// Run all tests
	scripttest.Test(t, context.Background(), engine, env, "testdata/*.txt")
}
