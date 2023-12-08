//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EC2Monitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateCreateCpuWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateCreateDiskOpsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateCreateDiskWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateCreateNetworkWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (e *jsiiProxy_EC2Monitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewEC2MonitoringParameters(scope MonitoringScope, props *EC2MonitoringProps) error {
	return nil
}

