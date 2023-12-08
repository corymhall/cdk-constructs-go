//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsClusterMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (r *jsiiProxy_RdsClusterMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (r *jsiiProxy_RdsClusterMonitoring) validateCreateConnectionsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (r *jsiiProxy_RdsClusterMonitoring) validateCreateCpuAndDiskUsageWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (r *jsiiProxy_RdsClusterMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (r *jsiiProxy_RdsClusterMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewRdsClusterMonitoringParameters(scope MonitoringScope, props *RdsClusterMonitoringProps) error {
	return nil
}

