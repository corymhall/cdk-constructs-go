//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoScalingGroupMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) validateCreateGroupSizeWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) validateCreateGroupStatusWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewAutoScalingGroupMonitoringParameters(scope MonitoringScope, props *AutoScalingGroupMonitoringProps) error {
	return nil
}

