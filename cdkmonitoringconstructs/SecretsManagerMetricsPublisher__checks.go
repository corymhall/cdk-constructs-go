//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

func (s *jsiiProxy_SecretsManagerMetricsPublisher) validateAddSecretParameters(secret awssecretsmanager.ISecret) error {
	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func validateSecretsManagerMetricsPublisher_GetInstanceParameters(scope MonitoringScope) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateSecretsManagerMetricsPublisher_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

