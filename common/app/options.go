package app

import (
	"github.com/veljkomatic/user-account/common/environment"
	"strings"
)

// Name returns an BuilderOption that sets the name of the app.
func Name(name ...string) BuilderOption {
	return func(builder *Builder) {
		appName := strings.Join(name, " ")
		if appName != "" {
			builder.appName = appName
		}
	}
}

// TaskDelegate returns an BuilderOption that adds a task to the app.
func TaskDelegate(taskName string, runFunc RunFunc) BuilderOption {
	return func(builder *Builder) {
		builder.addTask(taskName, runFunc)
	}
}

// VersionFromEnv returns an BuilderOption that sets the version of the app to the value of the environment variable.
// If env var is not set, the version is set to "0.0.0".
func VersionFromEnv(envAccessor environment.EnvAccessor) BuilderOption {
	return func(builder *Builder) {
		version := environment.ServiceVersion(envAccessor)
		builder.SetVersion(version)
	}
}
