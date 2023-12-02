package app

import (
	"github.com/veljkomatic/user-account/common/environment"
	"strings"
)

// Name returns an AppBuilderOption that sets the name of the app.
func Name(name ...string) AppBuilderOption {
	return func(builder *AppBuilder) {
		appName := strings.Join(name, " ")
		if appName != "" {
			builder.appName = appName
		}
	}
}

// TaskDelegate returns an AppBuilderOption that adds a task to the app.
func TaskDelegate(taskName string, runFunc RunFunc) AppBuilderOption {
	return func(builder *AppBuilder) {
		builder.addTask(taskName, runFunc)
	}
}

// VersionFromEnv returns an AppBuilderOption that sets the version of the app to the value of the environment variable.
// If envVar is not set, the version is set to "0.0.0".
func VersionFromEnv(envAccessor environment.EnvAccessor) AppBuilderOption {
	return func(builder *AppBuilder) {
		version := environment.ServiceVersion(envAccessor)
		builder.SetVersion(version)
	}
}
