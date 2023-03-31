//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

var env = map[string]string{
	"GOOS":        os.Getenv("GOOS"),
	"CGO_ENABLED": "0",
}

func Deps() error {
	err := sh.RunV("go", "mod", "download")
	if err != nil {
		return err
	}

	return sh.RunV("go", "mod", "tidy")
}

func UpdateDeps() error {
	return sh.RunV("go", "get", "-d", "-u", "./...")
}

func Lint() error {
	pkgs, err := listAllGoPackages()
	if err != nil {
		return fmt.Errorf("list packages: %w", err)
	}

	args := []string{"run"}
	args = append(args, pkgs...)

	return sh.RunV("golangci-lint", args...)
}

func listAllGoPackages() ([]string, error) {
	out, err := sh.Output("go", "list", "-f", "{{.Dir}}/...", "-m")
	if err != nil {
		return nil, err
	}

	return strings.Split(out, "\n"), nil
}

func listAllApps() ([]string, error) {
	out, err := sh.Output("go", "list", "-f", "{{.Dir}}", "-m")
	if err != nil {
		return nil, err
	}

	modules := make([]string, 0, 10)

	for _, path := range strings.Split(out, "\n")[1:] {
		modules = append(modules, filepath.Base(path))
	}

	return modules, nil
}
