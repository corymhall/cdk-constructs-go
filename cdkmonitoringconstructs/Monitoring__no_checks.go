//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Monitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (m *jsiiProxy_Monitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (m *jsiiProxy_Monitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewMonitoringParameters(scope MonitoringScope, props *BaseMonitoringProps) error {
	return nil
}

