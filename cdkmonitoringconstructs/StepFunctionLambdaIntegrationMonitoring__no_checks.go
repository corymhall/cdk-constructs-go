//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewStepFunctionLambdaIntegrationMonitoringParameters(scope MonitoringScope, props *StepFunctionLambdaIntegrationMonitoringProps) error {
	return nil
}

