package environment

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/veljkomatic/user-account/common/environment/deployment"
	"github.com/veljkomatic/user-account/common/ptr"
)

type Environment string

const (
	// EnvNotSet is used for things that must always be present, regardless of the env
	EnvNotSet      Environment = ""
	EnvProduction  Environment = "production"
	EnvDevelopment Environment = "development"
	EnvCanary      Environment = "canary"

	envExplicitlyNotSet = "explicitly_not_set"
)

func (e Environment) String() string {
	return string(e)
}

func GetReleaseEnvironment(serviceName, deployment string) Environment {
	if strings.HasPrefix(serviceName, "canary-") || strings.HasPrefix(deployment, "canary") {
		return EnvCanary
	}
	return EnvProduction

}

var (
	env                *Environment
	setEnvironmentOnce = sync.Once{}
)

const UserAccountEnvKey = "USER_ACCOUNT_ENV"

// EnvAccessor is an interface for accessing environment variables.
type EnvAccessor interface {
	GetEnv(key string) string
}

// RealEnvAccessor accesses real environment variables.
type RealEnvAccessor struct{}

// GetEnv retrieves the value of the environment variable named by the key.
func (r *RealEnvAccessor) GetEnv(key string) string {
	return os.Getenv(key)
}

func Get(envAccessor EnvAccessor) Environment {
	setEnvironmentOnce.Do(
		func() {
			if env != nil {
				return
			}
			envVar := envAccessor.GetEnv(UserAccountEnvKey)
			if envVar != "" {
				tmp := Environment(envVar)
				env = &tmp
			} else {
				env = ptr.From(EnvDevelopment)
			}
		},
	)

	return ptr.Value(env)
}

func Is(envAccessor EnvAccessor, e Environment) bool {
	return Get(envAccessor) == e
}

const (
	EnvServiceName    = "MY_SERVICE_NAME"
	EnvServiceVersion = "SERVICE_VERSION"
	EnvDeploymentName = "MY_DEPLOYMENT_NAME"
)

var (
	globalName                  = envExplicitlyNotSet
	globalServiceDeploymentName = envExplicitlyNotSet
)

// ServiceName returns name from environment or empty value if not set.
func ServiceName(envAccessor EnvAccessor) string {
	if globalName != envExplicitlyNotSet {
		return globalName
	}
	globalName = envAccessor.GetEnv(EnvServiceName)
	return globalName
}

func FullServiceName(envAccessor EnvAccessor) string {
	name := ServiceName(envAccessor)
	dpl := ServiceDeploymentName(envAccessor)
	if dpl == deployment.DeploymentNotSet || dpl == deployment.DeploymentEmpty ||
		dpl == deployment.DeploymentDefault {
		return name
	}

	return fmt.Sprintf("%s-%s", name, dpl)
}

func ServiceVersion(envAccessor EnvAccessor) string {
	version := envAccessor.GetEnv(EnvServiceVersion)
	if version == "" {
		version = "0.0.0"
	}

	return version
}

func ServiceDeploymentName(envAccessor EnvAccessor) deployment.DeploymentName {
	if globalServiceDeploymentName != envExplicitlyNotSet {
		return deployment.DeploymentName(globalServiceDeploymentName)
	}
	globalServiceDeploymentName = envAccessor.GetEnv(EnvDeploymentName)
	return deployment.DeploymentName(globalServiceDeploymentName)
}
