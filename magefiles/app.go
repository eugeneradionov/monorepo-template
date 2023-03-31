//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type App mg.Namespace

func (App) Deps(app string) error {
	if app == "all" {
		return inAllApps(Deps)
	}

	return inApp(app, Deps)
}

func (App) UpdateDeps(app string) error {
	if app == "all" {
		return inAllApps(UpdateDeps)
	}

	return inApp(app, UpdateDeps)
}

func (App) Lint(app string) error {
	if app == "all" {
		return inAllApps(Lint)
	}

	return inApp(app, Lint)
}

func (App) Build(app string) error {
	return sh.RunWithV(env, "go", "build", "-o", "./bin/"+app, "./"+app+"/cmd/")
}

func inApp(app string, fn func() error) error {
	if app == "" {
		return fn()
	}

	err := os.Chdir(app)
	if err != nil {
		return fmt.Errorf("cd %s: %w", app, err)
	}
	defer os.Chdir("./..")

	return fn()
}

func inApps(apps []string, fn func() error) error {
	for _, app := range apps {
		err := inApp(app, fn)
		if err != nil {
			return err
		}
	}

	return nil
}

func inAllApps(fn func() error) error {
	apps, err := listAllApps()
	if err != nil {
		return fmt.Errorf("lis all go modules: %w", err)
	}

	return inApps(apps, fn)
}
