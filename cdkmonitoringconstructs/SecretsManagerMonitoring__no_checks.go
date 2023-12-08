//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerMonitoring) validateCreateSecretsCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSecretsManagerMonitoringParameters(scope MonitoringScope, props *SecretsManagerMonitoringProps) error {
	return nil
}

