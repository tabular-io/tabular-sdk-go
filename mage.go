//go:build mage

package main

import (
	"fmt"
	"os/exec"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// Build generates tabular-sdk-go
func Build() error {
	fmt.Println("Building...")
	cmd := exec.Command(
		"openapi-generator",
		"generate",
		"-i",
		"openapispec.json",
		"-g",
		"go",
		"--package-name",
		"tabular",
		"--git-host",
		"github.com",
		"--git-user-id",
		"tabular-io",
		"--git-repo-id",
		"tabular-sdk-go",
		"-o",
		"tabular/")
	return cmd.Run()
}
