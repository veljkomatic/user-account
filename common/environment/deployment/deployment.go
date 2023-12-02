package deployment

type DeploymentName string

const (
	DeploymentNotSet  DeploymentName = "deployment_not_set"
	DeploymentEmpty   DeploymentName = ""
	DeploymentDefault DeploymentName = "default"
	DeploymentTest    DeploymentName = "test"
)

func (d DeploymentName) IsSet() bool {
	return d != DeploymentNotSet && d != DeploymentEmpty
}

func (d DeploymentName) String() string {
	return string(d)
}
