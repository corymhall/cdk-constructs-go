//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppSyncMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateCreateErrorCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateCreateErrorRateWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateCreateTpsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AppSyncMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewAppSyncMonitoringParameters(scope MonitoringScope, props *AppSyncMonitoringProps) error {
	return nil
}

