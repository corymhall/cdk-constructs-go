//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateCreateErrorCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateCreateErrorRateWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsCanaryMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSyntheticsCanaryMonitoringParameters(scope MonitoringScope, props *SyntheticsCanaryMonitoringProps) error {
	return nil
}

