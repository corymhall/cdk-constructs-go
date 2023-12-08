//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerMetricsPublisher) validateAddSecretParameters(secret awssecretsmanager.ISecret) error {
	return nil
}

func validateSecretsManagerMetricsPublisher_GetInstanceParameters(scope MonitoringScope) error {
	return nil
}

func validateSecretsManagerMetricsPublisher_IsConstructParameters(x interface{}) error {
	return nil
}

