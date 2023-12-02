package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/veljkomatic/user-account/common/environment/deployment"
	"github.com/veljkomatic/user-account/common/environment/mock"
)

func TestGetReleaseEnvironment(t *testing.T) {
	testCases := []struct {
		serviceName string
		deployment  string
		expectedEnv Environment
	}{
		{"canary-service", "", EnvCanary},
		{"prod-service", "canary", EnvCanary},
		{"service", "production", EnvProduction},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.serviceName+"_"+tc.deployment, func(t *testing.T) {
			env := GetReleaseEnvironment(tc.serviceName, tc.deployment)
			assert.Equal(t, tc.expectedEnv, env)
		})
	}
}

func TestGet_WithMock(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			UserAccountEnvKey: "production",
		},
	}
	assert.Equal(t, EnvProduction, Get(mockAccessor))

	// Should not change the environment
	/// because sync.Once is used
	mockAccessor.EnvVars[UserAccountEnvKey] = "development"
	assert.NotEqual(t, EnvDevelopment, Get(mockAccessor))
}

func TestIs_WithMock(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			UserAccountEnvKey: "production",
		},
	}
	assert.True(t, Is(mockAccessor, EnvProduction))
}

func TestServiceName(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			EnvServiceName: "my-service",
		},
	}
	assert.Equal(t, "my-service", ServiceName(mockAccessor))
}

func TestFullServiceName(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			EnvServiceName:    "my-service",
			EnvDeploymentName: "test-deployment",
		},
	}
	assert.Equal(t, "my-service-test-deployment", FullServiceName(mockAccessor))
}

func TestServiceVersion(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			EnvServiceVersion: "1.0.0",
		},
	}
	assert.Equal(t, "1.0.0", ServiceVersion(mockAccessor))

	// Test with the version environment variable unset
	mockAccessor.EnvVars[EnvServiceVersion] = ""
	assert.Equal(t, "0.0.0", ServiceVersion(mockAccessor))
}

func TestServiceDeploymentName(t *testing.T) {
	mockAccessor := &mock.EnvAccessor{
		EnvVars: map[string]string{
			EnvDeploymentName: "test-deployment",
		},
	}
	assert.Equal(t, deployment.DeploymentName("test-deployment"), ServiceDeploymentName(mockAccessor))
}

func TestDeploymentName_IsSet(t *testing.T) {
	testCases := []struct {
		deploymentName deployment.DeploymentName
		expected       bool
	}{
		{deployment.DeploymentTest, true},
		{deployment.DeploymentDefault, true},
		{deployment.DeploymentEmpty, false},
		{deployment.DeploymentNotSet, false},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(string(tc.deploymentName), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.deploymentName.IsSet())
		})
	}
}

func TestDeploymentName_String(t *testing.T) {
	testCases := []struct {
		deploymentName deployment.DeploymentName
		expected       string
	}{
		{deployment.DeploymentTest, "test"},
		{deployment.DeploymentDefault, "default"},
		{deployment.DeploymentEmpty, ""},
		{deployment.DeploymentNotSet, "deployment_not_set"},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(string(tc.deploymentName), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.deploymentName.String())
		})
	}
}
