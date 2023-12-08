//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (c *jsiiProxy_CustomMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (c *jsiiProxy_CustomMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewCustomMonitoringParameters(scope MonitoringScope, props *CustomMonitoringProps) error {
	return nil
}

