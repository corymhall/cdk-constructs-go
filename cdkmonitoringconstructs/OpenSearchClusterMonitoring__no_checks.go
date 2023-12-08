//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenSearchClusterMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewOpenSearchClusterMonitoringParameters(scope MonitoringScope, props *OpenSearchClusterMonitoringProps) error {
	return nil
}

