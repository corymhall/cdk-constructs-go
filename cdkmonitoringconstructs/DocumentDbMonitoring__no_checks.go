//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DocumentDbMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateConnectionsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateResourceUsageWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateTransactionsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewDocumentDbMonitoringParameters(scope MonitoringScope, props *DocumentDbMonitoringProps) error {
	return nil
}

