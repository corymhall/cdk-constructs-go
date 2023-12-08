//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (l *jsiiProxy_LogMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (l *jsiiProxy_LogMonitoring) validateCreateIncomingLogsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (l *jsiiProxy_LogMonitoring) validateResolveRecommendedHeightParameters(numRows *float64) error {
	return nil
}

func (l *jsiiProxy_LogMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewLogMonitoringParameters(scope MonitoringScope, props *LogMonitoringProps) error {
	return nil
}

