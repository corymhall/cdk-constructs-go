//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateCreateReadCapacityWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateCreateThrottlesWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateCreateWriteCapacityWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewDynamoTableGlobalSecondaryIndexMonitoringParameters(scope MonitoringScope, props *DynamoTableGlobalSecondaryIndexMonitoringProps) error {
	return nil
}

