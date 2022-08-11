package planner

import (
	"os"
	"path/filepath"
)

type GoPlanner struct{}

// GoPlanner implements interface Planner (compile-time check)
var _ Planner = (*GoPlanner)(nil)

func (g *GoPlanner) Name() string {
	return "GoPlanner"
}

func (g *GoPlanner) IsRelevant(srcDir string) bool {
	goModPath := filepath.Join(srcDir, "go.mod")
	return fileExists(goModPath)
}

func (g *GoPlanner) Plan(srcDir string) *BuildPlan {
	return &BuildPlan{
		Packages: []string{
			"go",
		},
		InstallCommand: "go get",
		BuildCommand:   "CGO_ENABLED=0 go build -o out",
		StartCommand:   "GIN_MODE=release ./out", // TODO: Move gin specific stuff elsewhere.
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}