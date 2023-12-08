//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewStepFunctionServiceIntegrationMonitoringParameters(scope MonitoringScope, props *StepFunctionServiceIntegrationMonitoringProps) error {
	return nil
}

