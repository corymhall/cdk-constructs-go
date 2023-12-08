//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionActivityMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewStepFunctionActivityMonitoringParameters(scope MonitoringScope, props *StepFunctionActivityMonitoringProps) error {
	return nil
}

