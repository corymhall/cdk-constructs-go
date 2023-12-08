//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerSecretMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSecretsManagerSecretMonitoringParameters(scope MonitoringScope, props *SecretsManagerSecretMonitoringProps) error {
	return nil
}

