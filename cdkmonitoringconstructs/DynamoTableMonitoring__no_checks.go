//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoTableMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateErrorsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateReadCapacityWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateThrottlesWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateCreateWriteCapacityWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateForEachOperationLatencyAlarmDefinitionParameters(operation awsdynamodb.Operation, alarm *map[string]*LatencyThreshold) error {
	return nil
}

func (d *jsiiProxy_DynamoTableMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewDynamoTableMonitoringParameters(scope MonitoringScope, props *DynamoTableMonitoringProps) error {
	return nil
}

