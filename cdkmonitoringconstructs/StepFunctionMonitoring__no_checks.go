//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_StepFunctionMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewStepFunctionMonitoringParameters(scope MonitoringScope, props *StepFunctionMonitoringProps) error {
	return nil
}

