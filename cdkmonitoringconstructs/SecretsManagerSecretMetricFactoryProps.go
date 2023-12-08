package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Experimental.
type SecretsManagerSecretMetricFactoryProps struct {
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
}

