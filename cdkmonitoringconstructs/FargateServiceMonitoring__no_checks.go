//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FargateServiceMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateCreateCpuWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateCreateMemoryWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateCreateTaskHealthWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateCreateTpcFlowsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (f *jsiiProxy_FargateServiceMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewFargateServiceMonitoringParameters(scope MonitoringScope, props *CustomFargateServiceMonitoringProps) error {
	return nil
}

