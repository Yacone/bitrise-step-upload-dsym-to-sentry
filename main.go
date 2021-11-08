package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/env"
)

type Config struct {
	OrgSlug     string `env:"sentry_org"`
	ProjectSlug string `env:"sentry_project"`
	AuthToken   string `env:"sentry_auth_token"`
}

func main() {
	pathDsyms := "./"
	var config Config
	envRepo := env.NewRepository()
	stepconf.NewInputParser(envRepo).Parse(&config)
	fmt.Println(config.AuthToken)

	cmdLog, err := exec.Command("sentry-cli", "--auth-token", config.AuthToken, "upload-dif", "--org", config.OrgSlug, "--project", config.ProjectSlug, pathDsyms).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to upload dsyms error: %#v | output: %s", err.Error(), cmdLog)
		os.Exit(-1)
	}

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
