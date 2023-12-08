package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

// Experimental.
type CodeBuildProjectMetricFactoryProps struct {
	// Experimental.
	Project awscodebuild.IProject `field:"required" json:"project" yaml:"project"`
}

