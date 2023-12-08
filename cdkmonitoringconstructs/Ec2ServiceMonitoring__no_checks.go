//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2ServiceMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateCreateCpuWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateCreateMemoryWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateCreateTaskHealthWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateCreateTpcFlowsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2ServiceMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewEc2ServiceMonitoringParameters(scope MonitoringScope, props *CustomEc2ServiceMonitoringProps) error {
	return nil
}

