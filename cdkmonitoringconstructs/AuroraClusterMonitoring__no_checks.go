//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuroraClusterMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateCreateConnectionsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateCreateCpuUsageWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateCreateServerlessDatabaseCapacityWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AuroraClusterMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewAuroraClusterMonitoringParameters(scope MonitoringScope, props *AuroraClusterMonitoringProps) error {
	return nil
}

